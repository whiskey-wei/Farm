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

//获取指定id的母牛的产犊记录
func GetCalve(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}

	if !valid.HasErrors() {
		if model.ExistCalveByID(id) {
			data = model.GetCalve(id)
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
	BirthTime := c.PostForm("birth_time")
	FlowingTime := c.PostForm("flowing_time")
	FetusTime := c.PostForm("fetus_time")
	FetusBirthTime := c.PostForm("fetus_birth_time")
	PlacentaTime := c.PostForm("placenta_time")

	code := e.INVALID_PARAMS
	valid := validation.Validation{}

	//判断时间格式
	var err error
	_, err = util.HandleTime(BirthTime)
	if err != nil {
		code = e.ERROR_TIME_FORMAT
		log.Println("birth_time")
	} else {
		_, err = util.HandleTime(FlowingTime)
		if err != nil {
			code = e.ERROR_TIME_FORMAT
			log.Println("flowing_time")
		} else {
			_, err = util.HandleTime(FetusTime)
			if err != nil {
				code = e.ERROR_TIME_FORMAT
				log.Println("fetus_time")
			} else {
				_, err = util.HandleTime(FetusBirthTime)
				if err != nil {
					code = e.ERROR_TIME_FORMAT
					log.Println("fetus_birth_time")
				} else {
					_, err = util.HandleTime(PlacentaTime)
					if err != nil {
						code = e.ERROR_TIME_FORMAT
						log.Println("placenta_time")
					}
				}
			}
		}
	}
	if code == e.INVALID_PARAMS {
		CowId := com.StrTo(c.PostForm("cow_id")).MustInt()
		FetusOrgan := c.PostForm("fetus_organ")
		IsComplete := com.StrTo(c.PostForm("is_complete")).MustInt()
		IsAbortion := com.StrTo(c.PostForm("is_abortion")).MustInt()
		YakId := com.StrTo(c.PostForm("yak_id")).MustInt()
		YakIndex := com.StrTo(c.PostForm("yak_index")).MustFloat64()
		MilkProduction := com.StrTo(c.PostForm("milk_production")).MustFloat64()
		Cream := com.StrTo(c.PostForm("cream")).MustFloat64()
		Protein := com.StrTo(c.PostForm("protein")).MustFloat64()

		//表单处理
		valid.Min(CowId, 1, "cow_id").Message("cow_id必须大于0")
		valid.Min(YakId, 1, "yak_id").Message("yal_id必须大于0")

		valid.Required(FetusOrgan, "fetus_oragan").Message("fetus_organ不能为空")

		valid.Range(IsComplete, 1, 2, "is_complete").Message("is_complete只能是1,2")
		valid.Range(IsAbortion, 1, 2, "is_abortion").Message("is_abortion只能是1,2")

		if !valid.HasErrors() && YakIndex > 0 && MilkProduction > 0 && Cream > 0 && Protein > 0 {
			if model.ExistCalveByID(CowId) {
				code = e.ERROR_EXIST
			} else {
				calve := model.CalveRecord{
					CowId:          CowId,
					FetusOrgan:     FetusOrgan,
					IsComplete:     IsComplete,
					IsAbortion:     IsAbortion,
					YakId:          YakId,
					YakIndex:       YakIndex,
					MilkProduction: MilkProduction,
					Cream:          Cream,
					Protein:        Protein,
					BirthTime:      BirthTime,
					FlowingTime:    FlowingTime,
					FetusTime:      FetusTime,
					FetusBirthTime: FetusBirthTime,
					PlacentaTime:   PlacentaTime,
				}
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

}

//删除产犊记录
func DeleteCalve(c *gin.Context) {

}
