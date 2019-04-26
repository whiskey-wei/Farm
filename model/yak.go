package model

type YakRecord struct { //生长发育记录
	Id                int    `gorm:"primary_key" json:"id"`
	YakId             int    `json:"yak_id"`             //犊牛号
	Variety           string `json:"variety"`            //品种
	BirthTime         string `json:"birth_time"`         //生日
	FatherNumber      int    `json:"father_number"`      //父亲号
	MotherNumber      int    `json:"mother_number"`      //母亲号
	Weight            int    `json:"weight"`             //体重
	Length            int    `json:"length"`             //体长
	Height            int    `json:"height"`             //体高
	Bust              int    `json:"bust"`               //胸围
	FrontrearDistance int    `json:"frontrear_distance"` //乳头前后间距
	LeftrightDistance int    `json:"leftright_distance"` //乳头左右间距
	LeftfrontLength   int    `json:"leftfront_length"`   //左前乳头长
	RightrearLength   int    `json:"rightrear_length"`   //右后乳头长
	RecordTime        int    `json:"record_time"`        //记录日期
}
