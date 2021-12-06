package views

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)


type WaterPumps struct {
	gorm.Model
	ID string `json:"ID" gorm:"primary_key"`
	TimeStamp int64 `json:"time"`
	Volume uint `json:"volume"`
	Temperature uint `json:"temperature"`
}

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open("mysql", "root:root1234@tcp(127.0.0.1:13306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {log.Fatal(err)}
	database.AutoMigrate(&WaterPumps{})

	DB = database
}

//GET
func GetData(c *gin.Context) {
	var getWaterpump []WaterPumps
	DB.Find(&getWaterpump)

	c.JSON(http.StatusOK, gin.H{"data": getWaterpump})
}

//POST
func CreateData(c *gin.Context) {
	var input WaterPumps

	//bindjson to bind recieved json to newData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//add it create slice
	//data := WaterPumps{}
	DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

//its only saving it one at a time but its not posting the new data
