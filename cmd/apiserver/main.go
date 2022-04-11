package main

import (
	"VideoHost/database"
	"VideoHost/routes"
	_ "github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	_ "html/template"
	"net/http"
	_ "net/http"
)

type error struct {
	Error string `json:"error"`
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Users)
}

func postUser(c *gin.Context) {

	//body := c.Request.GetBody
	var newUser database.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, error{"Bad request"})
		return
	}

	database.Users = append(database.Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func helloPage(c *gin.Context) {
	c.IndentedJSON(200, "Hello")
}

func main() {
	route := gin.Default()
	route.GET("/", helloPage)
	route.GET("/streams/info", routes.GetUserStreamKey)
	route.GET("/settings/stream_key", routes.GetSettings)
	route.POST("/settings/stream_key", routes.PostSettings)
	route.Run("localhost:3000")
}

/*
func hello_page(w http.ResponseWriter, r *http.Request) {
	bob := User{Name: "Bob", Nickname: "boba", Teamname: "Solit", Language: []string{"C#", "Java"}}
	//fmt.Fprintf(w, `<h1>Main</h1>
	//						<b>123</b>`)
	tmpl, _ := template.ParseFiles("templates/hello_page.html")
	tmpl.Execute(w, bob)
}
*/
