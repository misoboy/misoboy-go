package pagination

import (
	"strconv"
	"bytes"
	"fmt"
)

var (
	firstPageLabel = "<a href=\"#\" class=\"frist\" title=\"맨앞으로\" onclick=\"%s(%s); return false;\">&nbsp;</a>&#160;"
	previousPageLabel = "<a href=\"#\" class=\"prev\" title=\"이전\" onclick=\"%s(%s); return false;\">&nbsp;</a>&#160;"
	currentPageLabel = "<strong><a href=\"#\">%s</a></strong>&#160;"
	otherPageLabel = "<a href=\"#\" onclick=\"%s(%s); return false;\">%s</a>&#160;"
	nextPageLabel = "<a href=\"#\" class=\"next\" title=\"다음\" onclick=\"%s(%s); return false;\">&nbsp;</a>&#160;"
	lastPageLabel = "<a href=\"#\" class=\"last\" title=\"맨뒤로\" onclick=\"%s(%s); return false;\">&nbsp;</a>&#160;"
)

const PAGINE_ENABLE_ON  = "on"
const PAGINE_ENABLE_OFF = "off"

type Pagination struct {
	PaginationEnable 		string
	CurrentPageNo      		string
	RecordCountPerPageNo 	string
	PageSizeNo 			string
	CondOrder 			string
	CondAlign 			string
	TotalRecordCount		string
	JsFunction			string
}

func (c *Pagination) FirstPageNoOnPageList() string {
	currentPageNo, _ := strconv.ParseInt(c.CurrentPageNo, 10, 64)
	pageSizeNo, _ := strconv.ParseInt(c.PageSizeNo, 10, 64)
	return strconv.Itoa( int( ((currentPageNo - 1) / pageSizeNo) * pageSizeNo + 1 ) )
}

func (c *Pagination) LastPageNoOnPageList() string {
	firstPageNoOnPageList, _ := strconv.ParseInt(c.FirstPageNoOnPageList(), 10, 64)
	pageSizeNo, _ := strconv.ParseInt(c.PageSizeNo, 10, 64)
	lastPageNoOnPageList := firstPageNoOnPageList + pageSizeNo - 1
	totalPageCount, _ := strconv.ParseInt(c.TotalPageCount(), 10, 64)
	if lastPageNoOnPageList > totalPageCount {
		lastPageNoOnPageList = totalPageCount
	}
	return strconv.Itoa( int(lastPageNoOnPageList) )
}

func (c *Pagination) FirstRecordIndex() string {
	currentPage, _ := strconv.ParseInt(c.CurrentPageNo, 10, 64)
	recordCountPerPage, _ := strconv.ParseInt(c.RecordCountPerPageNo, 10, 64)
	firstRecordIndex := strconv.Itoa( int(( currentPage - 1 ) * recordCountPerPage) )
	return firstRecordIndex
}

func (c *Pagination) LastRecordIndex() string {
	currentPageNo, _ := strconv.ParseInt(c.CurrentPageNo, 10, 64)
	recordCountPerPageNo, _ := strconv.ParseInt(c.RecordCountPerPageNo, 10, 64)
	return strconv.Itoa(int( currentPageNo * recordCountPerPageNo ))
}

func (c *Pagination) PaginationSupport() map[string]string {
	pagination := make(map[string]string, 0)

	pagination["pagination"] = c.PaginationEnable
	pagination["pageIndex"] = c.CurrentPageNo
	pagination["recordCountPerPage"] = c.RecordCountPerPageNo
	pagination["pageSize"] = c.PageSizeNo

	return pagination
}

func (c *Pagination) TotalPageCount() string {
	totalRecordCount, _ := strconv.ParseInt(c.TotalRecordCount, 10, 64)
	recordCountPerPage, _ := strconv.ParseInt(c.RecordCountPerPageNo, 10, 64)
	totalPageCount := ( ( totalRecordCount - 1 ) / recordCountPerPage ) + 1
	return strconv.Itoa( int(totalPageCount) )
}

func (c *Pagination) FirstPageNo() string {
	return "1"
}

func (c *Pagination) LastPageNo() string {
	return c.TotalPageCount()
}

func (c *Pagination) RenderPagination() string {

	firstPageNo := c.FirstPageNo()
	firstPageNoOnPageList, _ := strconv.ParseInt(c.FirstPageNoOnPageList(), 10, 64)
	totalPageCount, _ := strconv.ParseInt(c.TotalPageCount(), 10, 64)
	pageSize, _ := strconv.ParseInt(c.PageSizeNo, 10, 64)
	lastPageNoOnPageList, _ := strconv.ParseInt(c.LastPageNoOnPageList(), 10, 64)
	currentPageNo, _ := strconv.ParseInt(c.CurrentPageNo, 10, 64)
	lastPageNo := c.LastPageNo()

	var buffer bytes.Buffer

	if totalPageCount > pageSize {
		if firstPageNoOnPageList > pageSize {
			buffer.WriteString(fmt.Sprintf(firstPageLabel, c.JsFunction, firstPageNo))
			buffer.WriteString(fmt.Sprintf(previousPageLabel, c.JsFunction, strconv.Itoa(int (firstPageNoOnPageList - 1) )))
		} else {
			buffer.WriteString(fmt.Sprintf(firstPageLabel, c.JsFunction, firstPageNo))
			buffer.WriteString(fmt.Sprintf(previousPageLabel, c.JsFunction, firstPageNo))
		}
	}

	for i := firstPageNoOnPageList; i <= lastPageNoOnPageList; i++ {
		if i == currentPageNo {
			buffer.WriteString(fmt.Sprintf(currentPageLabel, strconv.Itoa( int(i) )))
		} else {
			buffer.WriteString(fmt.Sprintf(otherPageLabel, c.JsFunction, strconv.Itoa( int(i) ), strconv.Itoa( int(i) ) ))
		}
	}

	if totalPageCount > pageSize {
		if lastPageNoOnPageList < totalPageCount {
			buffer.WriteString(fmt.Sprintf(nextPageLabel, c.JsFunction, strconv.Itoa( int(firstPageNoOnPageList + pageSize) )))
			buffer.WriteString(fmt.Sprintf(lastPageLabel, c.JsFunction, lastPageNo))
		} else {
			buffer.WriteString(fmt.Sprintf(nextPageLabel, c.JsFunction, lastPageNo))
			buffer.WriteString(fmt.Sprintf(lastPageLabel, c.JsFunction, lastPageNo))
		}
	}

	return buffer.String()
}

func (c *Pagination) SetTotalRecordCount(dataList []map[string]interface{}) {
	if dataList != nil && len(dataList) > 0 {
		dataMap := dataList[0]
		c.TotalRecordCount = dataMap["TOT"].(string)
		fmt.Println(c.TotalRecordCount)
	}
}