package trade

import "github.com/zhaoyoucai/aeop/top"

type ExtendsBuyerAcceptGoodsTimeAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ExtendsBuyerAcceptGoodsTimeAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.redefining.extendsbuyeracceptgoodstime")
	
	a.Ae = ae
	a.Api = api

}

func (a *ExtendsBuyerAcceptGoodsTimeAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *ExtendsBuyerAcceptGoodsTimeAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}