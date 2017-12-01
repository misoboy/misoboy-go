package bbs

import (
	_ "misoboy_web/models"
	bbsRepository "misoboy_web/repository/bbs"
)

type BbsService interface {
	GetList() []map[string] interface{}
	GetDetail(bbsId string, nttSn string) map[string] interface{}
}

type bbsService struct {
	bbsRepository bbsRepository.BbsRepository
	test string
}

func NewBbsService(bbsRepository bbsRepository.BbsRepository) BbsService {
	return &bbsService{ bbsRepository : bbsRepository }
}

/**
 * 게시글 목록
 */
func (s *bbsService) GetList() []map[string] interface{} {
	return s.bbsRepository.SelectBbsNttList()
	/*[]models.BbsVo{
		models.BbsVo{
			Title: "테스트1",
			Content: "테스트1 내용",
		},
		models.BbsVo{
			Title: "테스트2",
			Content: "테스트1 내용",
		},
		models.BbsVo{
			Title: "테스트3",
			Content: "테스트1 내용",
		},
		models.BbsVo{
			Title: "테스트4",
			Content: "테스트1 내용",
		},
		models.BbsVo{
			Title: "테스트5",
			Content: "테스트1 내용",
		},
	}*/
}

/**
 * 게시글 상세
 */
func (s *bbsService) GetDetail(bbsId string, nttSn string) map[string] interface{} {
	return s.bbsRepository.SelectBbsNttDetail(bbsId, nttSn)
}