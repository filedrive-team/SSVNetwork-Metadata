package models

type OperatorBaseInfo struct {
	Id             uint64 `gorm:"id;primaryKey" json:"operator_id"`
	Name           string `gorm:"name" json:"name"`
	OwnerAddress   string `gorm:"owner_address" json:"owner_address"`
	PublicKey      string `gorm:"public_key" json:"public_key"`
	Active         bool   `gorm:"active" json:"active"`
	Fee            string `gorm:"fee" json:"fee"`
	Score          uint32 `gorm:"score" json:"score"`
	IndexInOwner   uint32 `gorm:"index_in_owner" json:"index_in_owner"`
	ValidatorCount uint32 `gorm:"validator_count" json:"validator_count"`
	Timestamp      int64  `gorm:"timestamp" json:"timestamp"`
}

func init() {
	RegisterModel(&OperatorBaseInfo{})
}
