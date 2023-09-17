package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raulcoroiu/wowTeamComp/pkg/controllers"
)

//const localhost = ":8080"

func main() {

	r := gin.Default()

	r.GET("/getBestTeam", controllers.GetBestTeamHandler)
	r.Run(":8080")
}
