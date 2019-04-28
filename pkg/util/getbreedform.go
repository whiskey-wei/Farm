package util

import (
	"farm/model"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetBreedForm(c *gin.Context) (breedrecord model.BreedingRecord) {
	breedrecord.CowId = com.StrTo(c.PostForm("cow_id")).MustInt()
	breedrecord.StartTime = c.PostForm("start_time")
	breedrecord.EndTime = c.PostForm("end_time")
	breedrecord.FirstTime = c.PostForm("first_time")
	breedrecord.SecondTime = c.PostForm("second_time")
	breedrecord.ThirdTime = c.PostForm("third_time")
	breedrecord.FourthTime = c.PostForm("fourth_time")
	breedrecord.FinalTime = c.PostForm("final_time")

	breedrecord.FirstNumber = com.StrTo(c.PostForm("first_number")).MustInt()
	breedrecord.SecondNumber = com.StrTo(c.PostForm("second_number")).MustInt()
	breedrecord.ThirdNumber = com.StrTo(c.PostForm("third_number")).MustInt()
	breedrecord.FourthNumber = com.StrTo(c.PostForm("fourth_number")).MustInt()
	breedrecord.FinalNumber = com.StrTo(c.PostForm("final_number")).MustInt()

	return
}
