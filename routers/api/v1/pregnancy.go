/*
 * @Copyright @ 2019 Shopee
 * @Author : Lewei He
 * @Email : lewei.he@shopee.com
 * @Date : 2019-06-14 11:50
 */
package v1

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/whiskey-wei/Farm/model"
	"github.com/whiskey-wei/Farm/pkg/e"
	"github.com/whiskey-wei/Farm/pkg/setting"
	"github.com/whiskey-wei/Farm/pkg/util"
	"net/http"
	"time"
)

func GetPregnancys(c *gin.Context) {
	data := make(map[string]interface{})

	data["list"], _ = model.GetPregnancyRecords(util.GetPage(c), setting.PageSize)
	data["total"] = model.GetPregnancyRecordTotal()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

func GetPregnancy(c *gin.Context) {
	cowId := com.StrTo(c.Param("cow_id")).MustInt()
	data, err := model.GetPregnancyRecord(cowId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_NOT_EXIST,
			"msg":  e.GetMsg(e.ERROR_NOT_EXIST),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

func AddPregnancy(c *gin.Context) {
	var pregnancy model.PregnancyRecord
	//fmt.Println(c.PostForm("cow_id"))
	err := c.ShouldBind(&pregnancy)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	pregnancy.FinalTime, err = model.GetFinalBreedTime(pregnancy.CowId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	t, err := util.HandleTime(pregnancy.FinalTime)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	t = t.Add(time.Hour * 24 * 305)
	pregnancy.EstimateTime = util.HandleUnixTime(t)

	err = model.AddPregnancyRecord(pregnancy)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func UpdatePregnancy(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	var pregnancy model.PregnancyRecord
	//fmt.Println(c.PostForm("cow_id"))
	err := c.ShouldBind(&pregnancy)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	err = model.UpdatePregnancyRecord(id, pregnancy)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func DeletePregnancy(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	err := model.DeletePregnancyRecord(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
