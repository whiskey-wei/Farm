package model

import "fmt"

type CalveRecord struct {
	ID             string  //母牛号
	BirthTime      string  //分娩日期
	FlowingTime    string  //阴道开始流水日期
	FetusTime      string  //胎儿露出阴门时间
	FetusOrgan     string  //露出阴门的胎儿器官
	FetusBirthTime string  //胎儿娩出时间
	PlacentaTime   string  //胎盘排出时间
	IsComplete     bool    //胎衣排出是否完整
	IsAbortion     bool    //是否流产或者难产
	YakID          string  //犊牛编号
	YakIndex       float64 //犊牛相关指数
	MilkProduction float64 //泌乳期产奶量
	Cream          float64 //乳脂量
	Protein        float64 //乳蛋白量
} //产犊记录

func (this *CalveRecord) Add() {
	fmt.Println(this)
}

func (this *CalveRecord) Delete() {
	fmt.Println(this)
}

func (this *CalveRecord) Update() {
	fmt.Println(this)
}

func (this *CalveRecord) Select() {
	fmt.Println(this)
}
