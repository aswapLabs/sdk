package services;

import (
	"fmt"
	"strings"
	"context"
	"encoding/hex"
	Cli "github.com/starcoinorg/starcoin-go/client"
	"github.com/starcoinorg/starcoin-go/types"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"

	beego "github.com/beego/beego/v2/server/web"
)


const DEFAULT_MAX_GAS_AMOUNT = 10000000

func DoPairsRegister(tokenX, tokenY, network string) {

	ctx := context.Background()

	url := beego.AppConfig.DefaultString(
		network + "::rpc", 
		"https://barnard-seed.starcoin.org",
	) 
	
	client := Cli.NewStarcoinClient(url)


	adminAddr := beego.AppConfig.DefaultString(
		"admin_address", 
		"0x0A7B8DAb322448AF454FccAfFBCbF247",
	)
	
	

	privateKeyString := ""
	privateKeyBytes, _ := hex.DecodeString(privateKeyString)
	privateKey := types.Ed25519PrivateKey(privateKeyBytes)

	contractAddr, _ := types.ToAccountAddress(adminAddr)

	sender, _ := types.ToAccountAddress(adminAddr)


	moduleId := types.ModuleId{
		*contractAddr,
		"Router02",
	}

	tg1 := getTypeTag(tokenX)
	tg2 := getTypeTag(tokenY)

	scriptFunction := types.ScriptFunction{
		moduleId,
		"create_pair",
		[]types.TypeTag{tg1, tg2},
		[][]byte{},
	}

	payload := types.TransactionPayload__ScriptFunction{
		Value: scriptFunction,
	}


	price, err := client.GetGasUnitPrice(ctx)
	if err != nil {
	}

	state, err := client.GetState(ctx, adminAddr)

	if err != nil {
	}

	rawUserTransaction, err := client.BuildRawUserTransaction(ctx, *sender, &payload, price, DEFAULT_MAX_GAS_AMOUNT, state.SequenceNumber)
	if err != nil {
	}
	fmt.Printf("\n\n rawUserTransaction is : %v\n", rawUserTransaction)
	fmt.Printf("\n\n err is : %v\n", err)

	res, err := client.SubmitTransaction(ctx, privateKey,rawUserTransaction);

	fmt.Printf("\n hash is %#v\n\n", res)
	fmt.Printf("\n err is %#v\n\n", err)


	pendingTransactionInfo, err := client.GetPendingTransactionByHash(ctx, res);
	fmt.Printf("\n pendingTransactionInfo is %#v\n\n", pendingTransactionInfo)
	fmt.Printf("\n err is %#v\n\n", err)


	transactionInfo, err := client.GetTransactionInfoByHash(ctx, res);
	fmt.Printf("\ntransactionInfo is %#v\n\n", transactionInfo)
	fmt.Printf("\n err is %#v\n\n", err)

	transaction, err := client.GetTransactionByHash(ctx, res);
	fmt.Printf("\ntransactionInfo is %#v\n\n", transaction)
	fmt.Printf("\n err is %#v\n\n", err)

	
}

func getTypeTag(token string) *types.TypeTag__Struct {
	tokenArray := strings.Split(token, "::")

	tokenAddr, _ := types.ToAccountAddress(tokenArray[0])


	coinType := types.StructTag{
		Address: *tokenAddr,
		Module:  types.Identifier(tokenArray[1]),
		Name:    types.Identifier(tokenArray[2]),
	}

	return &types.TypeTag__Struct{Value: coinType}
}



func encode_u128_argument(arg serde.Uint128) []byte {

	s := bcs.NewSerializer()
	if err := s.SerializeU128(arg); err == nil {
		return s.GetBytes()
	}

	panic("Unable to serialize argument of type u64")
}



func encode_address_argument(arg types.AccountAddress) []byte {

	if val, err := arg.BcsSerialize(); err == nil {
		{
			return val
		}
	}

	panic("Unable to serialize argument of type address")
}


//
// type TypeTag__Address struct {
// }




