package model

import (
	"github.com/jinzhu/gorm"
)

type BreedingRecord struct {
	//配种记录
	Id            int    `gorm:"primary_key" json:"id" form:"id"`
	CowId         int    `json:"cow_id" form:"cow_id"`
	LastBirthTime string `json:"last_birth_time" form:"last_birth_time"` //上胎分娩时间
	StartTime     string `json:"start_time" form:"start_time"`           //发情开始时间
	EndTime       string `json:"end_time" form:"end_time"`               //发情结束时间
	FirstTime     string `json:"first_time" form:"first_time"`           //第一次配种时间
	FirstNumber   int    `json:"first_number" form:"first_number"`       //第一次配种冻精编号
	SecondTime    string `json:"second_time" form:"second_time"`
	SecondNumber  int    `json:"second_number" form:"second_number"`
	ThirdTime     string `json:"third_time" form:"third_time"`
	ThirdNumber   int    `json:"third_number" form:"third_number"`
	FourthTime    string `json:"fourth_time" form:"fourth_time"`
	FourthNumber  int    `json:"fourth_number" form:"fourth_number"`
	FinalTime     string `json:"final_time" form:"final_time"`
	FinalNumber   int    `json:"final_number" form:"final_number"`
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
		err = db.Where(maps).Limit(pageSize).Offset(pageNum).Find(&breedingrecords).Error
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

func AddBreedingRecord(breedingRecord BreedingRecord) bool {
	db.Create(&breedingRecord)
	return true
}

func UpdateBreedingRecord(id int, breedingRecord BreedingRecord) bool {
	return !db.Model(&BreedingRecord{}).Where("id = ?", id).Update(&breedingRecord).RecordNotFound()
}

func DeleteBreedigRecord(id int) error {
	return db.Where("id = ?", id).Delete(&BreedingRecord{}).Error
}

func GetFinalBreedTime(cowId int) (string, error) {
	var breed BreedingRecord
	err := db.Select("final_time").Where("cow_id = ?", cowId).Last(&breed).Error
	if err != nil {
		return "", err
	}
	return breed.FinalTime, nil
}

func GetFinalBreedNumber(cowId int) (int, error) {
	var breed BreedingRecord
	err := db.Select("final_number").Where("cow_id = ?", cowId).Last(&breed).Error
	return breed.FinalNumber, err
}
