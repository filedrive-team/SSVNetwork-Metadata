package repos

import (
	"fmt"
	"ssv_operator_metadata/db/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OperatorCid struct {
	Id  uint   `db:"id" json:"operator_id"`
	Cid string `db:"cid" json:"cid"`
}
type OperatorInfoRepo interface {
	GetOperatorInfoById(id uint64) (*models.OperatorInfo, error)
	GetOperatorInfos(offset, limit uint) (*[]models.OperatorInfo, error)
	GetOperatorCids(offset, limit uint) (*[]OperatorCid, error)
	GetOperatorAllRecordCids() (*map[string]string, error)
	UpdateOperatorInfo(info *models.OperatorInfo) (bool, error)
	UpdateCidRecordAndInfo(info *models.OperatorInfo, records []*models.CidRecord) (bool, error)
}

func NewOperatorInfo(db *gorm.DB) OperatorInfoRepo {
	return &ReposCli{
		db: db,
	}
}

func (dr *ReposCli) GetOperatorInfoById(id uint64) (*models.OperatorInfo, error) {
	info := models.OperatorInfo{}
	err := dr.db.First(&info, "id=?", id).Error
	if err == gorm.ErrRecordNotFound {
		// get base info
		var baseInfo models.OperatorBaseInfo
		if err := dr.db.First(&baseInfo, "id=?", id).Error; err != nil {
			return nil, err
		}
		info = models.OperatorInfo{
			OperatorSignMsg: models.OperatorSignMsg{
				Id:           baseInfo.Id,
				Name:         baseInfo.Name,
				OwnerAddress: baseInfo.OwnerAddress,
			},
		}
		return &info, nil
	}
	return &info, err
}

func (dr *ReposCli) GetOperatorInfos(offset, limit uint) (*[]models.OperatorInfo, error) {
	// todo : update
	info := []models.OperatorInfo{}
	err := dr.db.Raw(`select
	obi.id            as id,
	oi.name          as name,
	obi.owner_address as owner_address,
	ifnull(oi.location,'') as location,
	ifnull(oi.cloud_provider,'') as cloud_provider,
	ifnull(oi.eth1_node_client,'') as eth1_node_client,
	ifnull(oi.eth2_node_client,'') as eth2_node_client,
	ifnull(oi.description,'') as description,
	ifnull(oi.website_url,'') as website_url,
	ifnull(oi.twitter_url,'') as twitter_url,
	ifnull(oi.linkedin_url,'') as linkedin_url,
	ifnull(oi.discord_url,'') as discord_url,
	ifnull(oi.telegram_url,'') as telegram_url,
	ifnull(oi.relays_supported,'') as relays_supported,
	ifnull(oi.mev_bost_enabled,false) as mev_bost_enabled,
	ifnull(oi.logo,'') as logo,
	ifnull(oi.signature,'') as signature,
	ifnull(oi.timestamp,0) as timestamp,
	ifnull(oi.cid,'' ) as cid
	from operator_base_info obi
	left join operator_info oi on obi.id = oi.id
	order by oi.id desc, obi.id
limit  ? , ?`, offset, limit).Scan(&info).Error
	return &info, err
}

func (dr *ReposCli) GetOperatorCids(offset, limit uint) (*[]OperatorCid, error) {
	info := []OperatorCid{}
	err := dr.db.Model(models.OperatorInfo{}).
		Select("id", "cid").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Scan(&info).Error
	return &info, err
}

func (dr *ReposCli) GetOperatorAllRecordCids() (*map[string]string, error) {
	tmpLimt := 1000
	info := make(map[string]string)
	for offset, limit := uint(0), uint(tmpLimt); ; offset++ {
		tmp, err := dr.GetOperatorCids(offset, limit)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return &info, err
		}
		for _, v := range *tmp {
			info[fmt.Sprintf("%d", v.Id)] = v.Cid
		}

		if len(*tmp) < tmpLimt {
			break
		}
	}
	return &info, nil
}

func (dr *ReposCli) UpdateOperatorInfo(info *models.OperatorInfo) (bool, error) {
	if err := dr.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(info).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (dr *ReposCli) UpdateCidRecordAndInfo(info *models.OperatorInfo, records []*models.CidRecord) (bool, error) {
	tx := dr.db.Begin()

	// update cid-record
	for _, record := range records {
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "cid"}},
			UpdateAll: true,
		}).Create(record).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	// update operator info
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(info).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	// commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

func (dr *ReposCli) GetOperatorByKeyword(keyword string) (*[]models.OperatorInfo, error) {
	info := []models.OperatorInfo{}
	err := dr.db.Raw(`SELECT
	obi.id            as id,
	obi.name          as name,
	obi.owner_address as owner_address,
	ifnull(oi.location,'') as location,
	ifnull(oi.cloud_provider,'') as cloud_provider,
	ifnull(oi.eth1_node_client,'') as eth1_node_client,
	ifnull(oi.eth2_node_client,'') as eth2_node_client,
	ifnull(oi.description,'') as description,
	ifnull(oi.website_url,'') as website_url,
	ifnull(oi.twitter_url,'') as twitter_url,
	ifnull(oi.linkedin_url,'') as linkedin_url,
	ifnull(oi.discord_url,'') as discord_url,
	ifnull(oi.telegram_url,'') as telegram_url,
	ifnull(oi.logo,'') as logo,
	ifnull(oi.signature,'') as signature,
	ifnull(oi.timestamp,0) as timestamp,
	ifnull(oi.cid,'' ) as cid
	FROM operator_base_info obi
	left join operator_info oi on obi.id = oi.id
	WHERE CONCAT(obi.id,IFNULL(obi.name,'')) LIKE ?`, "%"+keyword+"%").Scan(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}
