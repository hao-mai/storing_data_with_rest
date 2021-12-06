package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hao-mai/sample_http_server/views"
)

func main() {

 router := gin.Default()

 views.ConnectDB()

 router.GET("/data", views.GetData)
 router.POST("/data", views.CreateData)
 router.Run("localhost:8000")

}
