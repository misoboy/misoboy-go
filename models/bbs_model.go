package models

import "time"

type BbsVo struct {
	BbsId string
	BbsNm string
	BbsIntrcn string
	UseAt string
	Ordr int64
	AccesAuthorAt string
	FrstRegistPnttm time.Time
	FrstRegisterId string
	LastUpdtPnttm time.Time
	LastUpdUsrId string
}