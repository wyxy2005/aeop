package logistics

import "github.com/zhaoyoucai/aeop/top"

type QuerySellerShipmentInfoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QuerySellerShipmentInfoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.aliexpress.logistics.querysellershipmentinfo")
	
	a.Ae = ae
	a.Api = api

}

func (a *QuerySellerShipmentInfoAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *QuerySellerShipmentInfoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}