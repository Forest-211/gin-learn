package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func middwareOne(c *gin.Context){
	fmt.Println("middleware one")
	c.Next()
	fmt.Println("middleware one end")
}
func middwareTwo(c *gin.Context){
	fmt.Println("middleware two")
	c.Next()
	fmt.Println("middleware two end")
}

func main() {
	r := gin.Default()

	r.GET("/", middwareOne, middwareTwo, func(c *gin.Context) {
		fmt.Println("request")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
			"data": map[string]string{
				"eamil": "78@qq.com",
			},
		})
	})
	/**
		gin中的中间件和koa中中间--洋葱模型差不多
		打印顺序：
			middleware one
			middleware two
			request
			middleware two end
			middleware one end
	*/

	r.Run(":9090")
}