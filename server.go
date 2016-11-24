package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
)

//DB stock the Database variable
var DB *neoism.Database

//APIGetAllLines return All lines of bus
func APIGetAllLines(c *gin.Context) {
	lines, err := GetAllLines()
	if err != nil {
		return
	}
	if value, ok := c.GetQuery("format"); ok == true && value == "json" {
		c.JSON(http.StatusOK, lines)
		return
	}
	content, err := json.Marshal(lines)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(content))
}

//APIGetStopsForLine return all stops for lineID
func APIGetStopsForLine(c *gin.Context) {
	lineID := c.Params.ByName("lineID")

	stops, err := GetStopsFromLineID(lineID)
	if err != nil {
		return
	}
	if value, ok := c.GetQuery("format"); ok == true && value == "json" {
		c.JSON(http.StatusOK, stops)
		return
	}
	content, err := json.Marshal(stops)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(content))
}

func init() {
	var err error

	protocol := os.Getenv("NEO4J_PROT")
	user := os.Getenv("NEO4J_USER")
	password := os.Getenv("NEO4J_PASSWD")
	host := os.Getenv("NEO4J_HOST")
	port := os.Getenv("NEO4J_PORT")
	DB, err = neoism.Connect(fmt.Sprintf("%s://%s:%s@%s:%s", protocol, user, password, host, port))
	if err != nil {
		return
	}
}

func main() {
	app := gin.Default()

	app.GET("/", APIGetAllLines)
	app.GET("/:lineID", APIGetStopsForLine)
	app.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "404", "message": "Page not found"})
	})
}
