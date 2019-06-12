package model

type DhiRecord struct { //DHI取样
	Id             int     `gorm:"primary_key" json:"id"`
	YakId          int     `json:"yak_id"`
	RecordTime     string  `json:"record_time"`
	BirthTime      string  `json:"birth_time"`
	DayAge         int     `json:"day_age"`
	Count          int     `json:"count"`
	MilkProduction float64 `json:"milk_production"`
}
