package main

import (
	"fmt"
	"money-transfer/handlers"
	"money-transfer/lib"

	"github.com/gin-gonic/gin"
)

const USER_COUNT = 2

func main() {

	createFiles()
	router := gin.Default()

	router.GET("/reset", handlers.Reset)
	router.POST("/transfer", handlers.Transfer)
	router.GET("/balances", handlers.GetBalance)

	router.Run(":5001")
}

func createFiles() (err error) {
	
	for i:=1; i <= USER_COUNT; i++ {
		err = lib.CreateFile(fmt.Sprintf("user%d.txt", i))
		if err != nil {
			return
		}
	}

	return
}