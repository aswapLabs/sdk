package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	services "aswap-go/services"
)

type PairsController struct {
	beego.Controller
}

type Result struct {
	Status int `json:"status"`
	Data string `json:"data"`
}

func (c *PairsController) Register() {
	tokenA := c.GetString("tokenA")
	tokenB := c.GetString("tokenB")
	chainId, _ := c.GetInt("chainId")

	res := Result{
		Status: 0,
		Data: "internal error",
	}

	if tokenA == tokenB {
		res.Data = "Token Can't Be Equal"	
		c.Data["json"] = res
    	c.ServeJSON()
    	return	
	}

	err := services.DoPairsRegister(tokenA, tokenB, chainId)

	fmt.Printf("\n the err is %v\n", err)

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