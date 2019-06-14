package model

import (
	"github.com/jinzhu/gorm"
)

type PregnancyRecord struct {
	//妊娠诊断
	Id                  int     `gorm:"primary_key" form:"id" path:"id"`
	CowId               int     `json:"cow_id" form:"cow_id"`
	CheckTime           string  `json:"check_time" form:"check_time"`                     //检查日期
	FinalTime           string  `json:"final_time" form:"final_time"`                     //配种日期
	EstimateTime        string  `json:"estimate_time" form:"estimate_time"`               //预计分娩日期
	OvaryDiameter       float64 `json:"ovary_diameter" form:"ovary_diameter"`             //卵巢直径
	FollicleDiameter    float64 `json:"follicle_diameter" form:"follicle_diameter"`       //卵泡直径
	UterineInflammation int     `json:"uterine_inflammation" form:"uterine_inflammation"` //子宫炎症
	VaginalInflammation int     `json:"vaginal_inflammation" form:"vaginal_inflammation"` //阴道炎症
	IsPregnancy         int     `json:"is_pregnancy" form:"is_pregnancy"`                 //是否妊娠
	IsEmpty             int     `json:"is_empty" form:"is_empty"`                         //是否空怀
}

func AddPregnancyRecord(pregnancy PregnancyRecord) error {
	return db.Create(&pregnancy).Error
}

func UpdatePregnancyRecord(id int, prenancy PregnancyRecord) error {
	return db.Model(PregnancyRecord{}).Where("id = ?", id).Update(&prenancy).Error
}

func GetPregnancyRecords(pageNum, pageSize int) (pregnancys []PregnancyRecord, err error) {
	//fmt.Println("pageNum:", pageNum, " pageSize:", pageSize)
	if pageNum > 0 && pageSize > 0 {
		//fmt.Println("exec this")
		err = db.Offset(pageNum).Limit(pageSize).Find(&pregnancys).Error
	} else {
		err = db.Find(&pregnancys).Error
	}
	if err != gorm.ErrRecordNotFound && err != nil {
		return nil, err
	}
	return pregnancys, nil
}

func GetPregnancyRecord(cowId int) (pregnancys []PregnancyRecord, err error) {
	err = db.Where("cow_id = ?", cowId).Find(&pregnancys).Error
	return
}

func DeletePregnancyRecord(id int) (err error) {
	err = db.Where("id = ?", id).Delete(&PregnancyRecord{}).Error
	return
}

func GetPregnancyRecordTotal() int {
	var count int
	db.Model(&PregnancyRecord{}).Count(&count)
	return count
}
