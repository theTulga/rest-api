package main

import (
	"fmt"
	"flag"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context){
	c.JSON(200, gin.H{ "message": "Hello World" })
}

func main() {
	prodFlagPtr := flag.Bool("prod", false, "a bool") 
	flag.Parse()
	if *prodFlagPtr {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("Starting in Production Mode")
	}

	r := gin.Default()

	r.GET("/", HomePage)
	r.Run()
}