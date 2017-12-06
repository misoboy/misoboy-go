package bbs

import (
	"sync"
	"misoboy_web/common/datasource"
)

type BbsRepository interface {
	SelectBbsNttList(params ...interface{}) []map[string]interface{}
	SelectBbsNttDetail(params ...interface{}) map[string]interface{}
	UpdateBbsNtt(params ...interface{}) int64
	DeleteBbsNtt(params ...interface{}) int64
}

func NewBbsRepository(dataSource datasource.DataSource) BbsRepository {
	return &sBbsRepository{ Datasource : dataSource }
}

type sBbsRepository struct {
	Datasource datasource.DataSource
	mu sync.RWMutex
}

func (r *sBbsRepository) SelectBbsNttList(params ...interface{}) []map[string]interface{} {
	return r.Datasource.SelectQuery("select *, (select count(1) from tb_bbs_ntt) as TOT from tb_bbs_ntt where bbs_id = ?", params...)
}

func (r *sBbsRepository) SelectBbsNttDetail(params ...interface{}) map[string]interface{} {
	return r.Datasource.SelectOneQuery("select * from tb_bbs_ntt where bbs_id = ? and ntt_sn = ?", params...)
}

func (r *sBbsRepository) UpdateBbsNtt(params ...interface{}) int64 {
	return r.Datasource.UpdateQuery("update tb_bbs_ntt set ntt_sj = ? where bbs_id = ?", params...)
}

func (r *sBbsRepository) DeleteBbsNtt(params ...interface{}) int64 {
	return r.Datasource.DeleteQuery("delete tb_bbs_ntt where bbs_id = ? and ntt_sn = ?", params...)
}