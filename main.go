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

	//controllers.PrintRuns(apiResponse)

	// Replace "monk_mistweaver" with the desired class specialization (e.g., "monk_mistweaver", "balance_druid", "protection_warrior")
	classSpec := "monk_mistweaver"

	topCompositions := controllers.FindTopCompositions(apiResponse, classSpec)
	for keystoneLevel, compositions := range topCompositions {
		fmt.Printf("Top 3 compositions for Keystone Level %d:\n", keystoneLevel)
		for i, compositionData := range compositions {
			fmt.Printf("%d. Composition: %v, Average Score: %.2f\n", i+1, compositionData.Composition, compositionData.AverageScore)
		}
		fmt.Println()
	}

}
