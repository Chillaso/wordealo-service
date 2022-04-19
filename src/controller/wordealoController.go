/*
	Create an endpoint to handle POST request on /wordealo containing a JSON body using Gin framework
	{
		"gameMode": 1,
		"wordLength": 5,
		"correct": [
			{"letter": "a", "position": 0},
			{"letter": "b", "position": -1}
		],
		"incorrect": [
			{"letter": "c", "position": 0}
		]
	}
*/
package controller

import (
	"log"
	"net/http"

	"github.com/chillaso/wordealo/model"
	"github.com/chillaso/wordealo/service"
	"github.com/gin-gonic/gin"
)

func GetWordle(context *gin.Context) {
	var wordealo *model.Wordealo
	err := context.Bind(&wordealo)
	if err == nil {
		result, err := service.GetPossibleWords(wordealo)
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusInternalServerError, err)
		}
		context.JSON(http.StatusOK, result)
	} else {
		log.Println(err)
		context.JSON(http.StatusBadRequest, err)
	}
}
