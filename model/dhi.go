package model

type DhiRecord struct {
	//DHI取样
	Id             int     `gorm:"primary_key" json:"id"`
	CowId          int     `json:"cow_id" form:"cow_id"`
	RecordTime     string  `json:"record_time" form:"record_time"`
	BirthTime      string  `json:"birth_time" form:"birth_time"`
	DayAge         int     `json:"day_age" form:"day_age"`
	Count          int     `json:"count" form:"count"`
	MilkProduction float64 `json:"milk_production" form:"milk_production"`
}

func AddDhiRecord(dhi DhiRecord) error {
	return db.Create(&dhi).Error
}

func UpdateDhiRecord(id int, dhi DhiRecord) error {
	return db.Model(&DhiRecord{}).Where("id = ?", id).Update(&dhi).Error
}

func DeleteDhiRecord(id int) error {
	return db.Where("id = ?", id).Delete(&DhiRecord{}).Error
}

func GetDhiTotal() int {
	var count int
	db.Model(&DhiRecord{}).Count(&count)
	return count
}

func GetDhiRecords(pageNum, pageSize int) ([]DhiRecord, error) {
	var dhiRecords []DhiRecord
	var err error
	if pageSize > 0 && pageNum > 0 {
		err = db.Offset(pageNum).Limit(pageSize).Find(&dhiRecords).Error
	} else {
		err = db.Find(&dhiRecords).Error
	}
	return dhiRecords, err
}

func GetDhiRecord(cowId int) ([]DhiRecord, error) {
	var dhiRecords []DhiRecord
	err := db.Where("cow_id = ?", cowId).Find(&dhiRecords).Error
	return dhiRecords, err
}
