package models

type CidRecord struct {
	Cid       string `gorm:"cid;primaryKey" json:"cid"`
	Data      string `gorm:"data" json:"data"`
	Note      string `gorm:"note;type:varchar(30)" json:"note"`
	IsSync    bool   `gorm:"is_sync" json:"is_sync"`
	CreatedAt int64  `gorm:"created_at;autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"updated_at;autoUpdateTime:milli" json:"updated_at"`
}

func init() {
	RegisterModel(&CidRecord{})
}

// func TableOfYear(user *interface{}, year int) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		tableName := db.Model(user).Name() + strconv.Itoa(year)
// 		return db.Table(tableName)
// 	}
// }
