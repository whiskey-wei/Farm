package v1

import (
	"farm/model"
	"farm/pkg/e"
	"farm/pkg/setting"
	"farm/pkg/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

//获取指定cowid的母牛的产犊记录
func GetCalve(c *gin.Context) {
	CowId := com.StrTo(c.Param("cow_id")).MustInt()

	valid := validation.Validation{}
	valid.Min(CowId, 1, "cow_id").Message("cow_id必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}

	if !valid.HasErrors() {
		if model.ExistCalveByCowID(CowId) {
			data = model.GetCalve(CowId)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST
		}
	} else {
		for _, err := range valid.Errors {
			log.Println("err.key:", err.Key, "err.message:", err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//获取多个母牛的产犊记录
func GetCalves(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		var err error
		data["lists"], err = model.GetCalves(util.GetPage(c), setting.PageSize, maps)
		if err != nil {
			code = e.ERROR
		}
		data["total"] = model.GetCalveTotal(maps)
		code = e.SUCCESS
	} else {
		for _, err := range valid.Errors {
			log.Println("err.key:", err.Key, " err.msg:", err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增产犊记录
func AddCalve(c *gin.Context) {

	//log.Println("add calve")
	calve := util.GetCalveForm(c)
	//log.Println("get map")
	code := e.INVALID_PARAMS

	//判断时间格式
	if !util.CheckTime(calve.FetusBirthTime) || !util.CheckTime(calve.FetusTime) ||
		!util.CheckTime(calve.FlowingTime) || !util.CheckTime(calve.PlacentaTime) {
		code = e.ERROR_TIME_FORMAT
	}
	if code == e.INVALID_PARAMS {
		valid := validation.Validation{}
		valid.Min(calve.CowId, 1, "cow_id").Message("cow_id应该大于0")
		valid.Min(calve.YakId, 1, "yak_id").Message("yak_id应该大于0")
		valid.Range(calve.IsComplete, 1, 2, "is_complete").Message("is_complete只能是1,2")
		valid.Range(calve.IsAbortion, 1, 2, "is_abortion").Message("is_abortion只能是1,2")
		valid.Required(calve.FetusOrgan, "fetus_organ").Message("fetus_organ")

		if !valid.HasErrors() && calve.YakIndex > 0 &&
			calve.MilkProduction > 0 && calve.Cream > 0 &&
			calve.Protein > 0 {
			if model.ExistCalveByID(calve.CowId) {
				code = e.ERROR_EXIST
			} else {
				model.AddCalve(calve)
				code = e.SUCCESS
			}
		} else {
			for _, err := range valid.Errors {
				log.Println("err.key:", err.Key, " err.msg:", err.Message)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//修改产犊记录
func UpdateCalve(c *gin.Context) {
	calve := util.GetCalveForm(c)
	calve.Id = com.StrTo(c.Param("id")).MustInt()
	log.Println(com.StrTo(c.Param("id")).MustInt())
	code := e.INVALID_PARAMS
	valid := validation.Validation{}
	if (calve.FetusBirthTime != "" && !util.CheckTime(calve.FetusBirthTime)) ||
		(calve.FetusTime != "" && !util.CheckTime(calve.FetusTime)) ||
		(calve.FlowingTime != "" && !util.CheckTime(calve.FlowingTime)) ||
		(calve.PlacentaTime != "" && !util.CheckTime(calve.PlacentaTime)) {
		code = e.ERROR_TIME_FORMAT
	} else {
		valid.Min(calve.Id, 1, "id").Message("id应该>0")
		if calve.YakId != 0 {
			valid.Min(calve.YakId, 1, "yak_id").Message("yakid应该大于0")
		}
		if calve.IsAbortion != 0 {
			valid.Range(calve.IsAbortion, 1, 2, "is_abortion").Message("is_abortion应该是1,2")
		}
		if calve.IsComplete != 0 {
			valid.Range(calve.IsComplete, 1, 2, "is_complete").Message("is_complete")
		}
		if !valid.HasErrors() {
			if !model.ExistCalveByID(calve.Id) {
				code = e.ERROR_NOT_EXIST
			} else {
				model.UpdateCalve(calve.Id, calve)
				code = e.SUCCESS
				log.Println("code:", code)
			}
		} else {
			for _, err := range valid.Errors {
				log.Println("err.key:", err.Key, " err.msg:", err.Message)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

//删除产犊记录
func DeleteCalve(c *gin.Context) {
	Id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS
	if model.ExistCalveByID(Id) {
		model.DeleteCalve(Id)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
