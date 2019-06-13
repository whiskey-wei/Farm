package v1

import (
	"net/http"

	"github.com/whiskey-wei/Farm/model"
	"github.com/whiskey-wei/Farm/pkg/e"
	"github.com/whiskey-wei/Farm/pkg/setting"
	"github.com/whiskey-wei/Farm/pkg/util"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// @Summary 获取配种记录
// @Description 通过cow_id获取
// @Accept  json
// @Produce  json
// @Param   cow_id     path    int     true        "母牛号"
// @Success 200 {string} string "ok"
// @Success 10002 {string} string "该数据不存在"
// @Router /api/v1/breeds/{cow_id} [get]
func GetBreed(c *gin.Context) {
	CowId := com.StrTo(c.Param("cow_id")).MustInt()
	var data interface{}
	var code int
	if model.ExistBreedingRecordByCowID(CowId) {
		data = model.GetBreedingRecord(CowId)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//获取多个母牛的配种记录
//
// @Summary 获取一页产犊记录
// @Description 每页20个记录，在路径中添加?page=x表示获取第x页的数据
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "页号"
// @Success 200 {string} string	"请求成功会返回id,cow_id,fetus_organ,is_complete,is_abortion,yak_id,yak_index,milk_production,cream,protein,birth_time,flowing_time,fetus_time,fetus_birth_time,placenta_time"
// @Success 500 {string} string "fail"
// @Failure 400 {object} string "请求参数错误"
// @Router /api/v1/breeds [get]

func GetBreeds(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	code := e.INVALID_PARAMS
	var ok bool

	data["list"], ok = model.GetBreedingRecords(util.GetPage(c), setting.PageSize, maps)
	if !ok {
		code = e.ERROR
	}
	data["total"] = model.GetBreedingRecordTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddBreed(c *gin.Context) {
	breed := util.GetBreedForm(c)
	code := e.INVALID_PARAMS

	if model.AddBreedingRecord(breed) {
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func DeleteBreed(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS
	if model.ExistBreedingRecordById(id) {
		if model.DeleteBreedigRecord(id) {
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func UpdateBreed(c *gin.Context) {
	breed := util.GetBreedForm(c)
	code := e.INVALID_PARAMS

	ok := model.UpdateBreedingRecord(breed.Id, breed)
	if ok {
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
