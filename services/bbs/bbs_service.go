package bbs

import (
	_ "misoboy_web/models"
	bbsRepository "misoboy_web/repository/bbs"
)

type BbsService interface {
	GetList(params ...interface{}) []map[string] interface{}
	GetDetail(bbsId string, nttSn string) map[string] interface{}
}

type bbsService struct {
	BbsRepository bbsRepository.BbsRepository
}

func NewBbsService(bbsRepository bbsRepository.BbsRepository) BbsService {
	return &bbsService{ BbsRepository : bbsRepository }
}

/**
 * 게시글 목록
 */
func (s *bbsService) GetList(params ...interface{}) []map[string] interface{} {
	return s.BbsRepository.SelectBbsNttList(params)
}

/**
 * 게시글 상세
 */
func (s *bbsService) GetDetail(bbsId string, nttSn string) map[string] interface{} {
	return s.BbsRepository.SelectBbsNttDetail(bbsId, nttSn)
}