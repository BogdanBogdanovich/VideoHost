package routes

import (
	"VideoHost/database"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
	_ "html/template"
	"net/http"
	_ "net/http"
)

func GetUserStreamKey(c *gin.Context) {
	var requestBody UserSettingRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, database.Error{Error: "Bad request"})
		return
	}
	var user = database.FindUserByUserName(requestBody.User.UserName)
	if user == nil {
		c.IndentedJSON(http.StatusBadRequest, database.Error{Error: "Пользователь не найден"})
		return
	}
	c.IndentedJSON(200, StreamKeyResponse{user.StreamKey})
}
