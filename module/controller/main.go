package controller

import (
	"time"
	"github.com/gin-gonic/gin"

	contract "github.com/michaelchandrag/warung-pintar/module/contract"
	usecase "github.com/michaelchandrag/warung-pintar/module/usecase"
	storage "github.com/michaelchandrag/warung-pintar/module/storage"
	websocket "github.com/michaelchandrag/warung-pintar/utils/websocket"
)

func ServePage (c *gin.Context) {
	c.HTML(200, "client.html", gin.H{
	})
}

func ServeWs (c *gin.Context) {
	websocket.ServeWs(websocket.WsHub, c.Writer, c.Request)
}

func Send (c *gin.Context) {
	type reqBody struct {
		Text 		string 		`json:"text"`
	}

	var payload reqBody
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Body request is not correct.",
		})
		return
	}

	currentTime := time.Now()
	message := storage.Message{
		Text: 		payload.Text,
		CreatedAt:	currentTime.Format("2006-01-02 15:04:05"),
	}

	var uc usecase.MessageUsecase

	var _message contract.MessageContract
	_message = message

	result, err := uc.Insert(_message)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Error when create message.",
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Success create message.",
		"data": result,
	})
	return
}

func Find (c *gin.Context) {
	var uc usecase.MessageUsecase
	var messages storage.Message
	var _messages contract.MessageContract = messages
	result, err := uc.Find(_messages)
	if err != nil {
		c.JSON(400, gin.H{
			"success": true,
			"message": "Something wrong when trying to get messages.",
			"error": err,	
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Success fetch messages.",
		"data": result,
	})
}