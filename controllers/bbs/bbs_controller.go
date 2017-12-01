package bbs

import (
	"github.com/kataras/iris/mvc"
	bbsService "misoboy_web/services/bbs"
	"fmt"
)

type BbsController struct {
	mvc.C
	service bbsService.BbsService
}

func (c *BbsController) GetMain() mvc.Result {
	return mvc.View{
		Layout: "layouts/layout.html",
		Name:   "bbs/main.html",
	}
}

func (c *BbsController) AnyList() mvc.Result {
	bbsList := c.service.GetList()

	return mvc.View{
		Layout: "layouts/layout.html",
		Name:   "bbs/list.html",
		Data: bbsList,
	}
}

func (c *BbsController) AnyDetail() mvc.Result {
	bbsId := c.Ctx.FormValue("bbsId")
	nttSn := c.Ctx.FormValue("nttSn")
	fmt.Println("bbsId : " + bbsId)
	fmt.Println("nttSn : " + nttSn)
	bbsDetail := c.service.GetDetail(bbsId, nttSn)
	return mvc.View{
		Layout: "layouts/layout.html",
		Name:   "bbs/detail.html",
		Data: bbsDetail,
	}
}