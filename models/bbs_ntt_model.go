package models

import "misoboy_web/common/pagination"

type BbsNttVo struct {
	Pagination *pagination.Pagination
	BbsVo BbsVo
	NttSn int64
	NttSj string
	NttCn string
	WrterId string
	WrterNm string
	AtchFileId string
	RdCnt int64
	BestFixingAt string
}

