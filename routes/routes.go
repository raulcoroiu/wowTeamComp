package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raulcoroiu/wowTeamComp/pkg/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/getBestTeam", controllers.GetBestTeamHandler)
	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginHandler)
}
