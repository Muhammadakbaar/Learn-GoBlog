package main

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"tampil": "sitamvan",
	})
}

func  main(){
	 r  := gin.Default()
	r.GET("/ping", ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}