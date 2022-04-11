package routes

import (
	"VideoHost/database"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "go/ast"
	_ "io/ioutil"
	_ "log"
	"net/http"
)

type StreamRequest struct {
	Streams database.Stream `json:"stream"`
}

func Streams(c *gin.Context) {

	var requestBody StreamRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, database.Error{Error: "Bad request"})
		return
	}

	for _, stream := range requestBody.Streams.StreamKey {

	}

	c.IndentedJSON(200, StreamKeyResponse{user.StreamKey})
}


if(req.query.streams){
let streams = JSON.parse(req.query.streams);
let query = {$or: []};
for (let stream in streams) {
if (!streams.hasOwnProperty(stream)) continue;
query.$or.push({stream_key : stream});
}

User.find(query,(err, users) => {
if (err)
return;
if (users) {
res.json(users);
}
});
}