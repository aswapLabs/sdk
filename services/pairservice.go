package services;

import (
	"fmt"
	"time"
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



func DoPairsRegister(tokenX, tokenY string, chainId int) error {

	ctx := context.Background()

	url := getNetworkByChainId(chainId)

	client := Cli.NewStarcoinClient(url)

	adminAddr := beego.AppConfig.DefaultString(
		"admin_address", 
		"0x580D0A84badec57956Fff0ddEdffa386",
	)

	b, _ := PairExists(tokenX, tokenY, adminAddr, ctx, client)
	if b == true {
		return nil
	}
	
	

	privateKeyString := ""
	privateKeyBytes, err := hex.DecodeString(privateKeyString)
	if err != nil {
		return err
	}

	privateKey := types.Ed25519PrivateKey(privateKeyBytes)

	contractAddr, err := types.ToAccountAddress(adminAddr)
	if err != nil {
		return err
	}

	sender, err := types.ToAccountAddress(adminAddr) 
	if err != nil {
		return err
	}

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
		return err
	}

	state, err := client.GetState(ctx, adminAddr)
	if err != nil {
		return err
	}

	rawUserTransaction, err := client.BuildRawUserTransaction(ctx, *sender, &payload, price, DEFAULT_MAX_GAS_AMOUNT, state.SequenceNumber)
	if err != nil {
		return err
	}

	res, err := client.SubmitTransaction(ctx, privateKey,rawUserTransaction);
	if err != nil {
		return err
	}

	var loops int
	for loops < 60 {
		pendingTransactionInfo, err := client.GetPendingTransactionByHash(ctx, res);
		loops ++;
		
		if pendingTransactionInfo.TransactionHash == "" || err != nil {
			txInfo, _ := client.GetTransactionInfoByHash(ctx, res)
			if(txInfo.TransactionHash == res) {
				break
			}
		} 
		
		time.Sleep(time.Second)
	}

	b, err = PairExists(tokenX, tokenY, adminAddr, ctx, client)
	if b == false || err != nil {
		return fmt.Errorf("not exist after")
	}

	return nil
	
}

func PairExists(
	tokenX, tokenY, adminAddr string, 
	ctx context.Context,
	client Cli.StarcoinClient,
) (bool, error) {
	call := Cli.ContractCall{
		adminAddr + "::Router02::pair_exists",
		[]string{
			tokenX, 
			tokenY,
		},
		[]string{},
	}

	callRes, err := client.CallContract(ctx, call)
	if err != nil {
		return false, err;
	}

	if !checkCallRes(callRes.([]interface{})) {
		return false, fmt.Errorf("not exist")
	}
	return true, nil
	
}

// 251 Barnard
// 254 Local
// 1 Main
func getNetworkByChainId(chainId int) string {

	var network string;

	switch chainId {
	case 1 :
		network = "Main"
	case 251: 
		network = "Barnard"
	case 254:
		network = "Local"
	default:
		network = "Barnard"
	}

	url := beego.AppConfig.DefaultString(
		network + "::rpc", 
		"https://barnard-seed.starcoin.org",
	) 
	return url
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


func checkCallRes(callRes []interface{}) bool { //1
	for _, val := range callRes {
		return val.(bool)
	}

	return false
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




