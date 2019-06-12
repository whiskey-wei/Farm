package model

import "github.com/jinzhu/gorm"

type PregnancyRecord struct { //妊娠诊断
	Id                  int     `gorm:"primary_key" json:"id"`
	CowId               int     `json:"cow_id"`
	CheckTime           string  `json:"check_time"`           //检查日期
	FinalTime           string  `json:"final_time"`           //配种日期
	EstimateTime        string  `json:"estimate_time"`        //预计分娩日期
	OvaryDiameter       float64 `json:"ovary_diameter"`       //卵巢直径
	FollicleDiameter    float64 `json:"follicle_diameter"`    //卵泡直径
	UterineInflammation int     `json:"uterine_inflammation"` //子宫炎症
	VaginalInflammation int     `json:"vaginal_inflammation"` //阴道炎症
	Pregnancy           int     `json:"pregnancy"`            //妊娠期
}

func AddPregnancyRecord(pregnancy PregnancyRecord) error {
	return db.Create(&pregnancy).Error
}

func UpdatePregnancyRecord(id int, prenancy PregnancyRecord) error {
	return db.Model(PregnancyRecord{}).Where("id = ?", id).Update(&prenancy).Error
}

func GetPregnancyRecords(pageNum, pageSize int) (pregnancys []PregnancyRecord, err error) {
	if pageNum > 0 && pageSize > 0 {
		err = db.Find(&pregnancys).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Find(&pregnancys).Error
	}
	if err != gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func GetPregnancyRecordById(cowId int) (pregnancys []PregnancyRecord, err error) {
	err = db.Where("cow_id = ?", cowId).Find(&pregnancys).Error
	return
}

func DeletePregnancyRecord(id int) (err error) {
	err = db.Where("id = ?", id).Delete(&PregnancyRecord{}).Error
	return
}
