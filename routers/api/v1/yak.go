/*
 * @Copyright @ 2019 Shopee
 * @Author : Lewei He
 * @Email : lewei.he@shopee.com
 * @Date : 2019-06-15 16:37
 */
package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/whiskey-wei/Farm/model"
	"github.com/whiskey-wei/Farm/pkg/setting"
	"github.com/whiskey-wei/Farm/pkg/util"
	"net/http"
	"time"
)

func GetYaks(c *gin.Context) {
	data := make(map[string]interface{})

	data["list"], _ = model.GetYakRecords(util.GetPage(c), setting.PageSize)
	data["total"] = model.GetYakRecordTotal()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

func GetYak(c *gin.Context) {
	yakId := com.StrTo(c.Param("yak_id")).MustInt()
	data, err := model.GetYakRecord(yakId)
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

func AddYak(c *gin.Context) {
	var yak model.YakRecord
	err := c.ShouldBind(&yak)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	//出生时间和母亲编号
	yak.BirthTime, yak.MotherNumber, err = model.GetBirthTimeAndCowIdByYakId(yak.YakId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	//父亲编号
	yak.FatherNumber, err = model.GetFinalBreedNumber(yak.MotherNumber)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	//记录时间
	yak.RecordTime = util.HandleUnixTime(time.Now())
	//fmt.Println(yak)
	err = model.AddYakRecod(yak)
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

func UpdataYak(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	var yak model.YakRecord
	err := c.ShouldBind(&yak)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	err = model.UpdateYakRecord(id, yak)
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

func DeleteYak(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	err := model.DeleteYakRecord(id)
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
