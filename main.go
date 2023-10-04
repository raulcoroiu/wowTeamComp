package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raulcoroiu/wowTeamComp/routes"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
