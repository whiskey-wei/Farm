/*
 * @Copyright @ 2019 Shopee
 * @Author : Lewei He
 * @Email : lewei.he@shopee.com
 * @Date : 2019-06-14 17:16
 */
package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/whiskey-wei/Farm/model"
	"github.com/whiskey-wei/Farm/pkg/e"
	"github.com/whiskey-wei/Farm/pkg/setting"
	"github.com/whiskey-wei/Farm/pkg/util"
	"net/http"
	"time"
)

func GetMilkProductions(c *gin.Context) {
	data := make(map[string]interface{})

	data["list"], _ = model.GetMilkProductionRecords(util.GetPage(c), setting.PageSize)
	data["total"] = model.GetMilkProductionRecordTotal()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

func GetMilkProduction(c *gin.Context) {
	cowId := com.StrTo(c.Param("cow_id")).MustInt()
	data, err := model.GetMilkProductionRecord(cowId)
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

func AddMilkProduction(c *gin.Context) {
	var milkProduction model.MilkProductionRecord
	err := c.ShouldBind(&milkProduction)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	milkProduction.SumProduction = milkProduction.MornProduction + milkProduction.NoonProduction + milkProduction.DuskProduction
	milkProduction.RecordTime = util.HandleUnixTime(time.Now())
	//fmt.Println(milkProduction)
	err = model.AddMilkProductionRecord(milkProduction)
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

func UpdateMilkProduction(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	var milkProduction model.MilkProductionRecord
	err := c.ShouldBind(&milkProduction)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	err = model.UpdateMilkProductionRecord(id, milkProduction)
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

func DeleteMilkProduction(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	err := model.DeleteMilkProductionRecord(id)
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
