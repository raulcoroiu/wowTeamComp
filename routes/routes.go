package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raulcoroiu/wowTeamComp/pkg/controllers"
)

func DefineRoutes(r *gin.Engine) {

	//r := gin.Default()+63
	r.GET("/getBestTeam", controllers.GetBestTeamHandler)

}
