/*
 * @Copyright @ 2019 Shopee
 * @Author : Lewei He
 * @Email : lewei.he@shopee.com
 * @Date : 2019-06-15 19:57
 */
package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/whiskey-wei/Farm/model"
	"github.com/whiskey-wei/Farm/pkg/setting"
	"github.com/whiskey-wei/Farm/pkg/util"
)

func GetDhis(c *gin.Context) {
	data := make(map[string]interface{})

	data["list"], _ = model.GetDhiRecords(util.GetPage(c), setting.PageSize)
	data["total"] = model.GetDhiTotal()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func GetDhi(c *gin.Context) {
	cowId := com.StrTo(c.Param("cow_id")).MustInt()
	data, err := model.GetDhiRecord(cowId)
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
		"data": data,
	})
}

func AddDhi(c *gin.Context) {
	var dhi model.DhiRecord
	err := c.ShouldBind(&dhi)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	dhi.BirthTime, _ = model.GetLastBirthTime(dhi.CowId)
	dhi.MilkProduction, err = model.GetSumMilkProductionByCowId(dhi.CowId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	birthTime, err := model.GetBirthTimeByYakId(dhi.CowId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	t1, err := util.HandleTime(birthTime)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	t2 := time.Now()
	t := t2.Sub(t1)
	dhi.DayAge = int(t.Hours() / 24)
	dhi.Count = model.GetCalveCountByCowId(dhi.CowId)
	fmt.Println("dhi:", dhi)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func UpdateDhi(c *gin.Context) {

}

func DeleteDhi(c *gin.Context) {

}
