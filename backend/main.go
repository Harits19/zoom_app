package main

import (
	"hello-world/access_token"
	"hello-world/env"
	"hello-world/meeting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	env.LoadEnv()

	err := access_token.InitAccessToken()

	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/meeting", func(c *gin.Context) {
		createMeetingRequest := meeting.CreateMeetingRequest{}

		if err := c.ShouldBind(&createMeetingRequest); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Wrong body request",
			})
			return
		}

		err := meeting.CreateMeeting(createMeetingRequest)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Success create meeting",
		})
	})

	r.PATCH("/meeting", func(c *gin.Context) {
		updateMeetingRequest := meeting.UpdateMeetingRequest{}

		if err := c.ShouldBind(&updateMeetingRequest); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Wrong body request",
			})
			return
		}

		err := meeting.UpdateMeeting(updateMeetingRequest)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Success update meeting",
		})
	})

	r.GET("/meeting", func(c *gin.Context) {

		response, err := meeting.GetAllMeeting()
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Success get all meeting",
			"data":    *response,
		})
	})
	r.Run()
}
