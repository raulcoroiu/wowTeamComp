package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/raulcoroiu/wowTeamComp/pkg/controllers"
)

func SetupRoutes(r *gin.Engine) {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}                   // Replace with your allowed origins (e.g., your frontend URL)
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Specify the allowed HTTP methods
	corsConfig.AllowHeaders = []string{"*"}                                       // Allow all headers
	corsConfig.AllowCredentials = true                                            // Allow cookies and credentials to be sent

	// Use the CORS middleware with the configuration
	r.Use(cors.New(corsConfig))

	r.GET("/getBestTeam", controllers.GetBestTeamHandler)
	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginHandler)
}
