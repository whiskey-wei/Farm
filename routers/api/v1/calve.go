package v1

import (
	"farm/model"
	"farm/pkg/e"
	"farm/pkg/setting"
	"farm/pkg/util"
	"log"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// @Summary 获取产犊记录
// @Description 通过cow_id获取
// @Accept  json
// @Produce  json
// @Param   cow_id     path    int     true        "母牛号"
// @Success 200 {string} string "ok"
// @Success 10002 {string} string "该数据不存在"
// @Router /api/v1/calves/{cow_id} [get]
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
//
// @Summary 获取一页产犊记录
// @Description 每页20个记录，在路径中添加?page=x表示获取第x页的数据
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "页号"
// @Success 200 {string} string	"请求成功会返回id,cow_id,fetus_organ,is_complete,is_abortion,yak_id,yak_index,milk_production,cream,protein,birth_time,flowing_time,fetus_time,fetus_birth_time,placenta_time"
// @Success 500 {string} string "fail"
// @Failure 400 {object} string "请求参数错误"
// @Router /api/v1/calves [get]
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
//
// @Summary 新增产犊记录
// @Description 新增产犊记录
// @Accept  json
// @Produce  json
// @Param  cow_id  body  int  true  "母牛号"
// @Param  yak_id  body  int  true  "犊牛号"
// @Param  is_complete  body  int  true  "胎衣排出是否完整,只能是1,2"
// @Param  is_abortion  body  int  true  "是否流产或者难产,只能是1,2"
// @Param  fetus_organ  body  string true  "露出阴门的胎儿器官"
// @Param  birth_time  body  string true "分娩日期 格式为2006-01-02 15:04:05 正则表达式严格匹配"
// @Param  flowint_time  body  string true  "阴道开始流水日期 格式为2006-01-02 15:04:05 正则表达式严格匹配"
// @Param  fetus_time  body  string true "胎儿露出阴门时间 格式为2006-01-02 15:04:05 正则表达式严格匹配"
// @Param fetus_birth_time body string true "胎儿娩出时间"
// @Param placenta_time body string true "胎盘排出时间"
// @Param yak_index body string true "犊牛相关指数,未确定计算方法,先任意输入"
// @Param milk_production body string true "泌乳期产奶量,未确定计算方法,先任意输入"
// @Param cream body string true "乳脂量,未确定计算方法,先任意输入"
// @Param protein body string true "乳蛋白量,未确定计算方法,先任意输入"
// @Success 200 {string} string	"ok"
// @Failure 10003 {string} string "时间格式出错"
// @Failure 10001 {string} string "该记录已经存在"
// @Failure 400 {string} string "请求参数错误"
// @Router /api/v1/calves [post]
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
//
// @Summary 更新数据
// @Description 通过id更新数据
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "这个id与cow_id不同，是数据库中的自增主键，在请求数据的时候，这个id会发送到前端,可在前端设置隐藏数据保存这个id"
// @Param  cow_id  body  int  false  "母牛号"
// @Param  yak_id  body  int  false  "犊牛号"
// @Param  is_complete  body  int  false  "胎衣排出是否完整,只能是1,2"
// @Param  is_abortion  body  int  false  "是否流产或者难产,只能是1,2"
// @Param  fetus_organ  body  string false  "露出阴门的胎儿器官"
// @Param  birth_time  body  string false "分娩日期 格式为2006-01-02 15:04:05 正则表达式严格匹配"
// @Param  flowint_time  body  string false  "阴道开始流水日期 格式为2006-01-02 15:04:05 正则表达式严格匹配"
// @Param  fetus_time  body  string false "胎儿露出阴门时间 格式为2006-01-02 15:04:05 正则表达式严格匹配"
// @Param fetus_birth_time body string false "胎儿娩出时间"
// @Param placenta_time body string false "胎盘排出时间"
// @Param yak_index body string false "犊牛相关指数,未确定计算方法,先任意输入"
// @Param milk_production body string false "泌乳期产奶量,未确定计算方法,先任意输入"
// @Param cream body string false "乳脂量,未确定计算方法,先任意输入"
// @Param protein body string false "乳蛋白量,未确定计算方法,先任意输入"
// @Success 200 {string} string	"ok"
// @Failure 10003 {string} string "时间格式出错"
// @Failure 10002 {string} string "该记录不存在"
// @Failure 400 {string} string "请求参数错误"
// @Router /api/v1/calves/{id} [put]
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

//
// @Summary 删除产犊记录
// @Description 通过id删除产犊记录
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "这个id与cow_id不同，是数据库中的自增主键，在请求数据的时候，这个id会发送到前端,可在前端设置隐藏数据保存这个id"
// @Success 200 {string} string	"ok"
// @Failure 10001 {string} string "该记录不存在"
// @Failure 400 {string} string "请求参数错误"
// @Router /api/v1/calves/{id} [delete]
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
