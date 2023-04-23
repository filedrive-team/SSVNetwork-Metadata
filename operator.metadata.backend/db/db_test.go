package db

import (
	"log"
	"os"
	"testing"

	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func TestDBGorm(t *testing.T) {
	gdbs := map[string]*gorm.DB{
		"mainnet": NewDBGorm(os.Getenv("TEST_DB_URL")),
		"testnet": NewDBGorm(os.Getenv("TEST_DB_TESENET_URL")),
	}
	testArgs := []string{
		"testnet",
		"mainnet",
	}

	for _, db_type := range testArgs {
		db, ok := gdbs[db_type]
		if !ok {
			log.Printf("get [%v] failed.\n", db_type)
			continue
		}
		var dbInfo string
		err := db.Clauses(dbresolver.Use(db_type)).Raw("SELECT DATABASE()").Scan(&dbInfo).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Fatalln(err)
			continue
		}
		log.Printf("[%v]%+v\n", db_type, dbInfo)
	}
}
