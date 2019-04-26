package model

type BreedingRecord struct {
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
