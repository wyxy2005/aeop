package issue

import "github.com/zhaoyoucai/aeop/top"

type ReverseLogisticUpdateAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ReverseLogisticUpdateAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.spain.issue.partner.rma.reverselogistic.state.update")
	
	a.Ae = ae
	a.Api = api

}

func (a *ReverseLogisticUpdateAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *ReverseLogisticUpdateAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}