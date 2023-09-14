package main

import (
	"fmt"

	"github.com/raulcoroiu/wowTeamComp/pkg/controllers"
)

func main() {
	responseBody, err := controllers.MakeRequest()
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	apiResponse, err := controllers.ParseResponse(responseBody)
	if err != nil {
		fmt.Println("Error parsing JSON resp", err)
		return
	}

	controllers.PrintRuns(apiResponse)

	//r := &routes.Router{}
	//http.ListenAndServe(":8000", r)

}
