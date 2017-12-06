package bbs

import (
	"github.com/kataras/iris/mvc"
	bbsService "misoboy_web/services/bbs"
	webPagination "misoboy_web/common/pagination"
	"github.com/misoboy/go-commons-lang/stringUtils"
)

type BbsController struct {
	mvc.C
	Service bbsService.BbsService
}

func (c *BbsController) GetMain() mvc.Result {
	return mvc.View{
		Name:   "bbs/main.html",
	}
}

func (c *BbsController) AnyListBy(bbsId string) mvc.Result {
	pagination := webPagination.Pagination{
		PaginationEnable : webPagination.PAGINE_ENABLE_ON,
		CurrentPageNo : stringUtils.DefaultString(c.Ctx.FormValue("pageIndex"), "1"),
		RecordCountPerPageNo : "10",
		PageSizeNo : "10",
		JsFunction: "DoBbs.refreshList",
	}

	bbsList := c.Service.GetList(bbsId, pagination.PaginationSupport())
	pagination.SetTotalRecordCount(bbsList)

	dataMap := make(map[string] interface{}, 0)
	dataMap["bbsList"] = bbsList
	dataMap["bbsId"] = bbsId
	dataMap["pageIndex"] = pagination.CurrentPageNo

	return mvc.View{
		Name:   "bbs/list.html",
		Data: map[string]interface{}{
			"dataMap" : dataMap,
			"pagination" : mvc.HTML(pagination.RenderPagination()),
		},

	}
}

func (c *BbsController) AnyDetail() mvc.Result {
	bbsId := c.Ctx.FormValue("bbsId")
	nttSn := c.Ctx.FormValue("nttSn")
	bbsDetail := c.Service.GetDetail(bbsId, nttSn)

	dataMap := make(map[string] interface{}, 0)
	dataMap["bbsDetail"] = bbsDetail
	dataMap["bbsId"] = bbsId

	return mvc.View{
		Name:   "bbs/detail.html",
		Data: dataMap,
	}
}