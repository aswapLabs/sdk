package controllers

import (
	"fmt"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	services "aswap-go/services"
)

type PairsController struct {
	beego.Controller
}

type ParisParams struct {
	TokenA string `json:"tokenA"`
	TokenB string `json:"tokenB"`
	ChainId int `json:"chainId"`
}

type Result struct {
	Status int `json:"status"`
	Data string `json:"data"`
}

func (c *PairsController) Register() {

	param := ParisParams{}

	res := Result{
		Status: 0,
		Data: "internal error",
	}

	err := json.Unmarshal(c.Ctx.Input.CopyBody(1<<32), &param)

	if err != nil {
		res.Data = "Params Errror"	
		c.Data["json"] = res
    	c.ServeJSON()
    	return	
	} 

	if param.TokenA == param.TokenB {
		res.Data = "Token Can't Be Equal"	
		c.Data["json"] = res
    	c.ServeJSON()
    	return	
	}

	err = services.DoPairsRegister(param.TokenA, param.TokenB, param.ChainId)
	fmt.Printf("\n err is %v\n", err)
	if err == nil {
		res.Data = "success"
		res.Status = 1
	} 

	c.Data["json"] = res
    c.ServeJSON()
}



// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }