package bbs

import (
	_ "misoboy_web/models"
	bbsRepository "misoboy_web/repository/bbs"
	"misoboy_web/models"
)

type BbsService interface {
	SelectBbsNttList(bbsNttVo models.BbsNttVo) []map[string] interface{}
	SelectBbsNttDetail(bbsNttVo models.BbsNttVo) map[string] interface{}
	SelectBbsNttKey(bbsNttVo models.BbsNttVo) map[string]interface{}
	UpdateBbsNtt(bbsNttVo models.BbsNttVo) int64
	DeleteBbsNtt(bbsNttVo models.BbsNttVo) int64
	InsertBbsNtt(bbsNttVo models.BbsNttVo) int64
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
func (s *bbsService) SelectBbsNttList(bbsNttVo models.BbsNttVo) []map[string] interface{} {
	return s.BbsRepository.SelectBbsNttList(bbsNttVo)
}

/**
 * 게시글 상세
 */
func (s *bbsService) SelectBbsNttDetail(bbsNttVo models.BbsNttVo) map[string] interface{} {
	return s.BbsRepository.SelectBbsNttDetail(bbsNttVo)
}

/**
 * 신규 키 생성
 */
func (r *bbsService) SelectBbsNttKey(bbsNttVo models.BbsNttVo) map[string]interface{} {
	return r.BbsRepository.SelectBbsNttKey(bbsNttVo)
}

/**
 * 게시글 수정
 */
func (s *bbsService) UpdateBbsNtt(bbsNttVo models.BbsNttVo) int64 {
	return s.BbsRepository.UpdateBbsNtt(bbsNttVo)
}

/**
 * 게시글 삭제
 */
func (s *bbsService) DeleteBbsNtt(bbsNttVo models.BbsNttVo) int64 {
	return s.BbsRepository.DeleteBbsNtt(bbsNttVo)
}

/**
 * 게시글 등록
 */
func (s *bbsService) InsertBbsNtt(bbsNttVo models.BbsNttVo) int64 {
	return s.BbsRepository.InsertBbsNtt(bbsNttVo)
}