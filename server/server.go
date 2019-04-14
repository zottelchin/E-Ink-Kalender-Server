package main

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ErcHRbrXh6aE7KCOfbuFzfvP6lxyoA", func(c *gin.Context) {
		content, err := ioutil.ReadFile("./cache.txt")
		if err != nil {
			c.String(500, stamp()+"Error; ............; ............; Datei konnte; .......; ........; nicht gelesen werden; ; ; ; ; ; ; ; ; ; ; ; ; ; ; ; ; ; ; ; ; ;")
		} else {
			c.String(200, string(content))
		}
	})
	r.Run(":6342")
}

func stamp() string {
	return strconv.Itoa(time.Now().Day()) + ". " + time.Now().Month().String() + " " + strconv.Itoa(time.Now().Year()) + ";"
}
