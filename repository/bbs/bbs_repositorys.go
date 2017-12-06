package bbs

import (
	"sync"
	"misoboy_web/common/datasource"
	"misoboy_web/models"
	"bytes"
)

type BbsRepository interface {
	SelectBbsNttList(bbsNttVo models.BbsNttVo) []map[string]interface{}
	SelectBbsNttDetail(bbsNttVo models.BbsNttVo) map[string]interface{}
	SelectBbsNttKey(bbsNttVo models.BbsNttVo) map[string]interface{}
	UpdateBbsNtt(bbsNttVo models.BbsNttVo) int64
	DeleteBbsNtt(bbsNttVo models.BbsNttVo) int64
	InsertBbsNtt(bbsNttVo models.BbsNttVo) int64
}

func NewBbsRepository(dataSource datasource.DataSource) BbsRepository {
	return &sBbsRepository{ Datasource : dataSource }
}

type sBbsRepository struct {
	Datasource datasource.DataSource
	mu sync.RWMutex
}

func (r *sBbsRepository) SelectBbsNttList(bbsNttVo models.BbsNttVo) []map[string]interface{} {
	return r.Datasource.SelectQuery("select *, (select count(1) from tb_bbs_ntt where bbs_id = #{bbsId}) as TOT from tb_bbs_ntt where bbs_id = #{bbsId}", bbsNttVo)
}

func (r *sBbsRepository) SelectBbsNttDetail(bbsNttVo models.BbsNttVo) map[string]interface{} {
	return r.Datasource.SelectOneQuery("select * from tb_bbs_ntt where bbs_id = #{bbsId} and ntt_sn = #{nttSn}", bbsNttVo)
}

func (r *sBbsRepository) SelectBbsNttKey(bbsNttVo models.BbsNttVo) map[string]interface{} {
	return r.Datasource.SelectOneQuery("select ifnull(max(ntt_sn), 0) + 1 as NTT_SN from tb_bbs_ntt where bbs_id = #{bbsId}", bbsNttVo)
}


func (r *sBbsRepository) UpdateBbsNtt(bbsNttVo models.BbsNttVo) int64 {
	return r.Datasource.UpdateQuery("update tb_bbs_ntt set ntt_cn = #{nttCn}, ntt_sj = #{nttSj} where bbs_id = #{bbsId} and ntt_sn = #{nttSn}", bbsNttVo)
}

func (r *sBbsRepository) DeleteBbsNtt(bbsNttVo models.BbsNttVo) int64 {
	return r.Datasource.DeleteQuery("delete from tb_bbs_ntt where bbs_id = #{bbsId} and ntt_sn = #{nttSn}", bbsNttVo)
}

func (r *sBbsRepository) InsertBbsNtt(bbsNttVo models.BbsNttVo) int64 {
	var buffer bytes.Buffer
	buffer.WriteString("insert into tb_bbs_ntt(bbs_id, ntt_sn, ntt_sj, ntt_cn, wrter_id, wrter_nm, rdcnt, best_fixing_at, use_at, frst_regist_pnttm, frst_register_id)")
	buffer.WriteString("values (#{bbsId}, #{nttSn}, #{nttSj}, #{nttCn}, #{wrterId}, #{wrterNm}, 0, 'N', 'Y', now(), 'admin')")
	return r.Datasource.UpdateQuery(buffer.String(), bbsNttVo)
}