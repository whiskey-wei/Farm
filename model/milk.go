package model

import "github.com/jinzhu/gorm"

type MilkProductionRecord struct {
	//产奶记录
	Id             int     `gorm:"primary_key" json:"id"` //默认主键
	CowId          int     `json:"cow_id"`
	RecordTime     int     `json:"record_time"`
	MornProduction float64 `json:"morn_production"` //早上产奶量
	NoonProduction float64 `json:"noon_production"`
	DuskProduction float64 `json:"dusk_production"`
	SumProduction  float64 `json:"sum_production"` //日产奶量
	MornInjection  float64 `json:"morn_injection"` //早上缩宫素注射量
	NoonInjection  float64 `json:"noon_injection"`
	DuskInjection  float64 `json:"dusk_injection"`
}

func AddMilkProductionRecord(milkProduction MilkProductionRecord) error {
	return db.Create(&milkProduction).Error
}

func UpdateMilkProductionRecord(id int, milkProduction MilkProductionRecord) error {
	return db.Model(&MilkProductionRecord{}).Where("id = ?", id).Update(milkProdution).Error
}

func GetMilkProductionRecords(pageNum, pageSize int) ([]MilkProductionRecord, error) {
	var milkProductions []MilkProductionRecord
	var err error
	if pageSize > 0 && pageNum > 0 {
		err = db.Offset(pageNum).Limit(pageSize).Find(&milkProductions).Error
	} else {
		err = db.Find(&milkProductions).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return milkProductions, nil
}

func GetMilkProductionRecord(cowId int) ([]MilkProductionRecord, error) {
	var milkProduction []MilkProductionRecord
	err := db.Where("cow_id = ?", cowId).Find(&milkProduction).Error
	return milkProduction, err
}

func DeleteMilkProductionRecord(id int) error {
	return db.Where("id = ?", id).Delete(&MilkProductionRecord{}).Error
}
