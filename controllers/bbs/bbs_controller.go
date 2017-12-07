package bbs

import (
	"github.com/kataras/iris/mvc"
	bbsService "misoboy_web/services/bbs"
	webPagination "misoboy_web/common/pagination"
	"github.com/misoboy/go-commons-lang/stringUtils"
	"misoboy_web/models"
	"strconv"
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

/**
 * 게시글 목록
 * Any : /bbs/list/{bbsId}
 */
func (c *BbsController) AnyListBy(bbsId string) mvc.Result {
	pagination := webPagination.Pagination{
		PaginationEnable : webPagination.PAGINE_ENABLE_ON,
		CurrentPageNo : stringUtils.DefaultString(c.Ctx.FormValue("pageIndex"), "1"),
		RecordCountPerPageNo : "10",
		PageSizeNo : "10",
		JsFunction: "DoBbs.refreshList",
	}
	// pagination.PaginationSupport()
	bbsVo := models.BbsVo{BbsId : bbsId}
	bbsNttVo := models.BbsNttVo{ BbsVo : bbsVo, Pagination: &pagination }
	bbsList := c.Service.SelectBbsNttList(bbsNttVo)
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

/**
 * 게시글 조회
 * Get : /bbs/detail/{bbsId}/{nttSn}
 */
func (c *BbsController) GetDetailBy(bbsId string, nttSn int64) mvc.Result {
	bbsVo := models.BbsVo{BbsId : bbsId}
	bbsNttVo := models.BbsNttVo{ BbsVo : bbsVo, NttSn: nttSn }

	bbsDetail := c.Service.SelectBbsNttDetail(bbsNttVo)

	dataMap := make(map[string] interface{}, 0)
	dataMap["bbsDetail"] = bbsDetail
	dataMap["bbsId"] = bbsId

	return mvc.View{
		Name:   "bbs/detail.html",
		Data: map[string]interface{}{
			"dataMap" : dataMap,
		},
	}
}

/**
 * 게시글 수정 화면
 * Get : /bbs/update/{bbsId}/{nttSn}
 */
func (c *BbsController) GetUpdateBy(bbsId string, nttSn int64) mvc.Result {
	bbsVo := models.BbsVo{BbsId : bbsId}
	bbsNttVo := models.BbsNttVo{ BbsVo : bbsVo, NttSn: nttSn }

	bbsDetail := c.Service.SelectBbsNttDetail(bbsNttVo)

	dataMap := make(map[string] interface{}, 0)
	dataMap["bbsDetail"] = bbsDetail
	dataMap["bbsId"] = bbsId

	return mvc.View{
		Name:   "bbs/update.html",
		Data: map[string]interface{}{
			"dataMap" : dataMap,
		},
	}
}

/**
 * 게시글 수정
 * Put : /bbs/update/{bbsId}/{nttSn}
 */
func (c *BbsController) PutUpdateBy(bbsId string, nttSn int64) interface {} {
	nttSj := c.Ctx.FormValue("nttSj")
	nttCn := c.Ctx.FormValue("nttCn")
	bbsVo := models.BbsVo{BbsId : bbsId}
	bbsNttVo := models.BbsNttVo{ BbsVo : bbsVo, NttSn: nttSn, NttSj: nttSj, NttCn: nttCn }

	rs := c.Service.UpdateBbsNtt(bbsNttVo)

	return map[string]interface{}{
		"result" : rs,
	}
}

/**
 * 게시글 삭제
 * Delete : /bbs/delete/{bbsId}/{nttSn}
 */
func (c *BbsController) DeleteDeleteBy(bbsId string, nttSn int64) interface {} {
	bbsVo := models.BbsVo{BbsId : bbsId}
	bbsNttVo := models.BbsNttVo{ BbsVo : bbsVo, NttSn: nttSn }

	rs := c.Service.DeleteBbsNtt(bbsNttVo)

	return map[string]interface{}{
		"result" : rs,
	}
}

/**
 * 게시글 등록 화면
 * Get : /bbs/insert/{bbsId}
 */
func (c *BbsController) GetInsertBy(bbsId string) interface {} {

	dataMap := make(map[string] interface{}, 0)
	dataMap["bbsId"] = bbsId

	return mvc.View{
		Name:   "bbs/insert.html",
		Data: map[string]interface{}{
			"dataMap" : dataMap,
		},
	}
}

/**
 * 게시글 등록
 * Post : /bbs/insert/{bbsId}
 */
func (c *BbsController) PostInsertBy(bbsId string) interface {} {
	bbsVo := models.BbsVo{BbsId : bbsId}
	nttSj := c.Ctx.FormValue("nttSj")
	nttCn := c.Ctx.FormValue("nttCn")
	bbsNttVo := models.BbsNttVo{ BbsVo : bbsVo, NttSj: nttSj, NttCn: nttCn, WrterId: "admin", WrterNm: "관리자" }

	bbsNttKey := c.Service.SelectBbsNttKey(bbsNttVo)
	nttSn, _ := strconv.ParseInt(bbsNttKey["NTT_SN"].(string), 10, 64)
	bbsNttVo.NttSn = nttSn

	rs := c.Service.InsertBbsNtt(bbsNttVo)

	return map[string]interface{}{
		"result" : rs,
		"nttSn" : bbsNttVo.NttSn,
	}
}