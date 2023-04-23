package repos

import (
	"ssv_operator_metadata/db/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CidRecordRepo interface {
	GetOperatorsIndexCid() (string, error)
	GetCidRecord(cid string) (*models.CidRecord, error)
	GetNotSyncCids() (*[]models.CidRecord, error)
	UpdateCidRecord(info *models.CidRecord) (bool, error)
}

func NewCidRecordRepo(db *gorm.DB) CidRecordRepo {
	return &ReposCli{
		db: db,
	}
}

func (dr *ReposCli) GetOperatorsIndexCid() (string, error) {
	indexCid := ""
	err := dr.db.Model(models.CidRecord{}).
		Select("cid").
		Where("note='index'").
		Order("created_at desc").
		Limit(1).
		Scan(&indexCid).
		Error
	return indexCid, err
}

func (dr *ReposCli) GetCidRecord(cid string) (*models.CidRecord, error) {
	info := models.CidRecord{}
	err := dr.db.First(&info, "cid=?", cid).Error
	return &info, err
}

func (dr *ReposCli) GetNotSyncCids() (*[]models.CidRecord, error) {
	info := []models.CidRecord{}
	err := dr.db.
		Model(models.CidRecord{}).
		Where("is_sync = 0").
		Order("created_at asc").
		Limit(10).
		Scan(&info).
		Error
	return &info, err
}

func (dr *ReposCli) UpdateCidRecord(info *models.CidRecord) (bool, error) {
	if err := dr.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "cid"}},
		UpdateAll: true,
	}).Create(info).Error; err != nil {
		return false, err
	}
	return true, nil
}
