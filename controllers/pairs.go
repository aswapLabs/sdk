package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	services "aswap-go/services"
)

type PairsController struct {
	beego.Controller
}

func (c *PairsController) Register() {
	tokenX := c.GetString("token_x")
	tokenY := c.GetString("token_y")
	network := c.GetString("network")
	err := services.DoPairsRegister(tokenX, tokenY, network)
	var status bool

	if err == nil {
		status = true
	}

	c.Data["json"] = status
    c.ServeJSON()
}



// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }