package repos

import (
	"ssv_operator_metadata/db/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewOperatorBaseInfoRepo(db *gorm.DB) OperatorBaseInfoRepo {
	return &ReposCli{
		db: db,
	}
}

type OperatorBaseInfoRepo interface {
	GetOperatorBaseInfoById(id uint64) (*models.OperatorBaseInfo, error)
	OperatorsBaseInfoTotal() (uint, error)
	UpdateOperatorBaseInfo(info *models.OperatorBaseInfo) (bool, error)
}

func (dr *ReposCli) GetOperatorBaseInfoById(id uint64) (*models.OperatorBaseInfo, error) {
	info := models.OperatorBaseInfo{}
	err := dr.db.First(&info, "id=?", id).Error
	return &info, err
}
func (dr *ReposCli) OperatorsBaseInfoTotal() (uint, error) {
	total := int64(0)
	err := dr.db.Model(models.OperatorBaseInfo{}).Count(&total).Error
	return uint(total), err
}

func (dr *ReposCli) UpdateOperatorBaseInfo(info *models.OperatorBaseInfo) (bool, error) {
	if err := dr.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(info).Error; err != nil {
		return false, err
	}
	return true, nil
}
