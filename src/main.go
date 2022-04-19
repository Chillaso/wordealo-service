package main

import (
	"log"

	"github.com/chillaso/wordealo/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controller.RegisterController(router)
	router.Run(":8081")
	log.Println("Server is running on port 8081")
}
