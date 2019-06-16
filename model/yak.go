package model

type YakRecord struct {
	//生长发育记录
	Id                int     `gorm:"primary_key" json:"id" form:"id"`
	YakId             int     `json:"yak_id" form:"yak_id"`                           //犊牛号
	Variety           string  `json:"variety" form:"variety"`                         //品种
	BirthTime         string  `json:"birth_time" form:"birth_time"`                   //生日
	FatherNumber      int     `json:"father_number" form:"father_number"`             //父亲号
	MotherNumber      int     `json:"mother_number" form:"mother_number"`             //母亲号
	Weight            float64 `json:"weight" form:"weight"`                           //体重
	Length            float64 `json:"length" form:"length"`                           //体长
	Height            float64 `json:"height" form:"height"`                           //体高
	Bust              float64 `json:"bust" form:"bust"`                               //胸围
	FrontRearDistance float64 `json:"front_rear_distance" form:"front_rear_distance"` //乳头前后间距
	LeftRightDistance float64 `json:"left_right_distance" form:"left_right_distance"` //乳头左右间距
	LeftFrontLength   float64 `json:"left_front_length" form:"left_front_length"`     //左前乳头长
	RightRearLength   float64 `json:"right_rear_length" form:"right_rear_length"`     //右后乳头长
	RecordTime        string  `json:"record_time" form:"record_time"`                 //记录日期
}

func AddYakRecod(yakRecord YakRecord) error {
	return db.Create(&yakRecord).Error
}

func UpdateYakRecord(id int, yakRecord YakRecord) error {
	return db.Model(&YakRecord{}).Where("id = ?", id).Update(&yakRecord).Error
}

func DeleteYakRecord(id int) error {
	return db.Where("id = ?", id).Delete(&YakRecord{}).Error
}

func GetYakRecordTotal() int {
	var count int
	db.Model(&YakRecord{}).Count(&count)
	return count
}

func GetYakRecords(pageNum, pageSize int) ([]YakRecord, error) {
	var yakRecords []YakRecord
	var err error
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset(pageNum).Limit(pageSize).Find(&yakRecords).Error
	} else {
		err = db.Find(&yakRecords).Error
	}
	return yakRecords, err
}

func GetYakRecord(yakId int) ([]YakRecord, error) {
	var yakRecord []YakRecord
	err := db.Where("yak_id = ?", yakId).Find(&yakRecord).Error
	return yakRecord, err
}

func GetBirthTimeByYakId(yakId int) (string, error) {
	var yak YakRecord
	err := db.Select("birth_time").Where("yak_id = ?", yakId).First(&yak).Error
	return yak.BirthTime, err
}
