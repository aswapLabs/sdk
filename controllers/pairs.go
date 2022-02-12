package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	services "aswap-go/services"
)

type PairsController struct {
	beego.Controller
}

func (c *PairsController) Register() {
	// argsTypeTokenX := c.GetString("token_x")
	// argsTypeTokenY := c.GetString("token_y")
	services.DoPairsRegister()
}



// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }