package models

var autoMigrateModels []interface{}

func RegisterModel(m interface{}) {
	autoMigrateModels = append(autoMigrateModels, m)
}

func DbModels() []interface{} {
	return autoMigrateModels
}
