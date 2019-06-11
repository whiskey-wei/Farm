package model

import (
	"github.com/jinzhu/gorm"
)

type BreedingRecord struct { //配种记录
	Id            int    `gorm:"primary_key" json:"id"`
	CowId         int    `json:"cow_id"`
	LastBirthTime string `json:"last_birth_time"` //上胎分娩时间
	StartTime     string `json:"start_time"`      //发情开始时间
	EndTime       string `json:"end_time"`        //发情结束时间
	FirstTime     string `json:"first_time"`      //第一次配种时间
	FirstNumber   int    `json:"first_number"`    //第一次配种冻精编号
	SecondTime    string `json:"second_time"`
	SecondNumber  int    `json:"second_number"`
	ThirdTime     string `json:"third_time"`
	ThirdNumber   int    `json:"third_number"`
	FourthTime    string `json:"fourth_time"`
	FourthNumber  int    `json:"fourth_number"`
	FinalTime     string `json:"final_time"`
	FinalNumber   int    `json:"final_number"`
}

func ExistBreedingRecordByCowID(CowId int) bool {
	var ans int
	db.Select("id").Where("cow_id = ?", CowId).First(&ans)
	if ans > 0 {
		return true
	}
	return false
}

func ExistBreedingRecordById(Id int) bool {
	var ans int
	db.Select("id").Where("id = ?", Id).First(&ans)
	if ans > 0 {
		return true
	}
	return false
}

func GetBreedingRecordTotal(maps interface{}) (count int) {
	db.Model(&BreedingRecord{}).Where(maps).Count(&count)
	return
}

func GetBreedingRecords(pageNum, pageSize int, maps interface{}) (breedingrecords []BreedingRecord, flag bool) {
	var err error
	flag = true
	if pageNum > 0 && pageSize > 0 {
		err = db.Where(maps).Find(&breedingrecords).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&breedingrecords).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		flag = false
		breedingrecords = nil
	}
	return
}

func GetBreedingRecord(CowId int) (breedrecords []BreedingRecord) {
	db.Where("cow_id = ?", CowId).Find(&breedrecords)
	return
}

func AddBreedingRecord(breedingrecord BreedingRecord) bool {
	db.Create(&breedingrecord)
	return true
}

func UpdateBreedingRecord(id int, breedingrecord BreedingRecord) bool {
	return !db.Model(&BreedingRecord{}).Where("id = ?", id).Update(&breedingrecord).RecordNotFound()
}

func DeleteBreedigRecord(id int) bool {
	db.Where("id = ?", id).Delete(&BreedingRecord{})
	return true
}
