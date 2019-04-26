package model

type MilkproductionRecord struct {
	Id             int     `gorm:"primary_key" json:"id"` //默认主键
	CowId          int     `json:"cow_id"`
	RecordTime     int     `json:"record_time"`
	MornProduction float64 `json:"morn_production"`	//早上产奶量
	NoonProduction float64 `json:"noon_production"`
	DuskProduction float64 `json:"dusk_production"`
	SumProduction  float64 `json:"sum_production"`	//日产奶量
	MornInjection  float64 `json:"morn_injection"`	//早上缩宫素注射量
	NoonInjection  float64 `json:"noon_injection"`
	DuskInjection  float64 `json:"dusk_injection"`
}
