package marketing

import "github.com/zhaoyoucai/aeop/top"

type QueryPromotionProductlistAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryPromotionProductlistAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.redefining.queryproductlistresultforpromotion")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryPromotionProductlistAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *QueryPromotionProductlistAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}