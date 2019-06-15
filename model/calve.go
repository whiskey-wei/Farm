package model

import (
	"github.com/jinzhu/gorm"
)

type CalveRecord struct {
	//产犊记录
	Id             int     `gorm:"primary_key" json:"id"`
	CowId          int     `json:"cow_id"`           //母牛号
	FetusOrgan     string  `json:"fetus_organ"`      //露出阴门的胎儿器官
	IsComplete     int     `json:"is_complete"`      //胎衣排出是否完整
	IsAbortion     int     `json:"is_abortion"`      //是否流产或者难产
	YakId          int     `json:"yak_id"`           //犊牛编号
	YakIndex       float64 `json:"yak_index"`        //犊牛相关指数
	MilkProduction float64 `json:"milk_production"`  //泌乳期产奶量
	Cream          float64 `json:"cream"`            //乳脂量
	Protein        float64 `json:"protein"`          //乳蛋白量
	BirthTime      string  `json:"birth_time"`       //分娩日期
	FlowingTime    string  `json:"flowing_time"`     //阴道开始流水日期
	FetusTime      string  `json:"fetus_time"`       //胎儿露出阴门时间
	FetusBirthTime string  `json:"fetus_birth_time"` //胎儿娩出时间
	PlacentaTime   string  `json:"placenta_time"`    //胎盘排出时间
} //产犊记录

func ExistCalveByCowID(id int) bool {
	var calverecord CalveRecord
	db.Select("cow_id").Where("cow_id = ?", id).First(&calverecord)

	if calverecord.CowId > 0 {
		return true
	}

	return false
}

func ExistCalveByID(id int) bool {
	var calve CalveRecord
	db.Select("id").Where("id = ?", id).First(&calve)
	if calve.Id > 0 {
		return true
	}
	return false
}

func GetCalveTotal(maps interface{}) (count int) {
	db.Model(&CalveRecord{}).Where(maps).Count(&count)
	return
}

func GetCalves(pageNum int, pageSize int, maps interface{}) (calves []CalveRecord, err error) {
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&calves).Error
	} else {
		err = db.Where(maps).Find(&calves).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		calves = nil
	}
	return
}

func GetCalve(CowId int) (calves []CalveRecord) {
	db.Where("cow_id = ?", CowId).Find(&calves)
	return
}

func AddCalve(calve CalveRecord) bool {
	db.Create(&calve)
	return true
}

func UpdateCalve(id int, calve CalveRecord) bool {
	db.Model(&CalveRecord{}).Where("id = ?", id).Update(calve)
	return true
}

func DeleteCalve(id int) bool {
	db.Where("id = ?", id).Delete(&CalveRecord{})
	return true
}

func GetLastBirthTime(CowId int) (birthtime string, flag bool) {
	var calve CalveRecord
	flag = !db.Select("birth_time").Where("cow_id = ?", CowId).Last(&calve).RecordNotFound()
	birthtime = calve.BirthTime
	return
}

func GetBirthTimeAndCowIdByYakId(yakId int) (string, int, error) {
	var calve CalveRecord
	err := db.Select("birth_time, cow_id").Where("yak_id = ?", yakId).First(&calve).Error
	return calve.BirthTime, calve.CowId, err
}

func GetCowIdByYakId(yakId int) (int, error) {
	var calve CalveRecord
	err := db.Select("cow_id").Where("yak_id = ?", yakId).First(&calve).Error
	return calve.CowId, err
}
