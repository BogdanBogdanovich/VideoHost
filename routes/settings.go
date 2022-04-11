package routes

import (
	"VideoHost/database"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "html/template"
	"net/http"
	_ "net/http"
)

type UserSettingRequest struct {
	User database.User `json:"user"`
}

type StreamKeyResponse struct {
	StreamKey string `json:"stream_key"`
}

func GetSettings(c *gin.Context) {
	var requestBody UserSettingRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, database.Error{Error: "Bad request"})
		return
	}
	var user = database.FindUserByEmail(requestBody.User.Email)
	if user == nil {
		c.IndentedJSON(http.StatusBadRequest, database.Error{Error: "Пользователь не найден"})
		return
	}
	c.IndentedJSON(200, StreamKeyResponse{user.StreamKey})
}

func PostSettings(c *gin.Context) {
	var requestBody UserSettingRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, database.Error{Error: "Bad request"})
		return
	}
	var user = database.FindUserByEmail(requestBody.User.Email)
	if user == nil {
		c.IndentedJSON(http.StatusBadRequest, database.Error{Error: "Пользователь не найден"})
		return
	}
	user.StreamKey = GenerateStreamKey()
	database.SaveUser(*user)
	c.IndentedJSON(200, StreamKeyResponse{user.StreamKey})
}

func GenerateStreamKey() string {
	id := uuid.New()
	return id.String()
}
