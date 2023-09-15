package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raulcoroiu/wowTeamComp/pkg/controllers"
)

func DefineRoutes(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/getBestTeam", controllers.GetBestTeamHandler)

}
