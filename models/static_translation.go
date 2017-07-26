package models

type StaticTranslation struct {
	Base

	TextAlias         string ` gorm:"type:text"`
	StaticTranslation string `gorm:"type:text"`
}
