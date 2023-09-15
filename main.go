package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raulcoroiu/wowTeamComp/pkg/controllers"
)

func main() {
	//responseBody, err := controllers.MakeRequest()
	//if err != nil {
	//	fmt.Println("Error making request:", err)
	//	return
	//}

	//apiResponse, err := controllers.ParseResponse(responseBody)
	//if err != nil {
	//	fmt.Println("Error parsing JSON resp", err)
	//	return
	//}

	r := gin.Default()

	// Define your API routes
	r.GET("/getBestTeam", controllers.GetBestTeamHandler)

	// Start the Gin server on port 8080.
	r.Run(":8080")
}
