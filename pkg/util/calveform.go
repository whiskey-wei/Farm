package util

import (
	"farm/model"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetCalveForm(c *gin.Context) (calve model.CalveRecord) {
	calve.BirthTime = c.PostForm("birth_time")
	calve.FlowingTime = c.PostForm("flowing_time")
	calve.FetusTime = c.PostForm("fetus_time")
	calve.FetusBirthTime = c.PostForm("fetus_birth_time")
	calve.PlacentaTime = c.PostForm("placenta_time")
	calve.CowId = com.StrTo(c.PostForm("cow_id")).MustInt()
	calve.FetusOrgan = c.PostForm("fetus_organ")
	calve.IsComplete = com.StrTo(c.PostForm("is_complete")).MustInt()
	calve.IsAbortion = com.StrTo(c.PostForm("is_abortion")).MustInt()
	calve.YakId = com.StrTo(c.PostForm("yak_id")).MustInt()
	calve.YakIndex = com.StrTo(c.PostForm("yak_index")).MustFloat64()
	calve.MilkProduction = com.StrTo(c.PostForm("milk_production")).MustFloat64()
	calve.Cream = com.StrTo(c.PostForm("cream")).MustFloat64()
	calve.Protein = com.StrTo(c.PostForm("protein")).MustFloat64()
	return
}
