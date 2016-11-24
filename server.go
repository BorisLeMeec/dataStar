package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
)

//DB stock the Database variable
var DB *neoism.Database

//APIGetAllLines return All lines of bus
func APIGetAllLines(c *gin.Context) {
	content, err := GetAllLines()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	if value, ok := c.GetQuery("format"); ok == true && value == "json" {
		c.JSON(200, content)
		return
	}
}

//APIGetStopsForLine return all stops for lineID
func APIGetStopsForLine(c *gin.Context) {
	lineID := c.Params.ByName("lineID")

	content, err := GetStopsFromLineID(lineID)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	c.JSON(200, content)

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
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})
}
