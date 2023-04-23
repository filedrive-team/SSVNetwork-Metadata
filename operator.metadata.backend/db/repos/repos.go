package repos

import "gorm.io/gorm"

type ReposCli struct {
	db *gorm.DB
}

func NewRepoCli(db *gorm.DB) *ReposCli {
	return &ReposCli{
		db: db,
	}
}
