package controllers

import ()

type LotteryController struct {
	BaseController
}

func (this *LotteryController) Get() {
	this.Data["IsLottery"] = true
	this.TplNames = "lottery.html"
}
