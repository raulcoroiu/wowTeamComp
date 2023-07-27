package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/raulcoroiu/wowTeamComp/pkg/models"
)

const (
	apiURL = "https://raider.io/api/v1/mythic-plus/runs?season=season-df-1&region=world&dungeon=all&page=0"
)

// makeRequest sends an HTTP GET request to the API endpoint and returns the response body.
func MakeRequest() ([]byte, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func ParseResponse(responseBody []byte) (*models.ApiResponse, error) {
	var apiResponse models.ApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

func PrintRun(run models.Run) {
	fmt.Println("Run:")
	fmt.Printf("Season: %s\n", run.Season)
	fmt.Printf("Status: %s\n", run.Status)
	fmt.Printf("Keystone Run ID: %d\n", run.KeystoneRunID)
	fmt.Printf("Mythic Level: %d\n", run.MythicLevel)
	fmt.Printf("Clear Time (ms): %d\n", run.ClearTimeMS)
	fmt.Printf("Keystone Time (ms): %d\n", run.KeystoneTimeMS)
	fmt.Printf("Completed At: %s\n", run.CompletedAt)
	fmt.Printf("Number of Chests: %d\n", run.NumChests)
	fmt.Printf("Time Remaining (ms): %d\n", run.TimeRemainingMS)

	// mai trebuie adaugate
}

func PrintRuns(apiResponse *models.ApiResponse) {
	fmt.Println("Mythic+ Runs:")

	for i, ranking := range apiResponse.Rankings {
		fmt.Printf("Ranking %d:\n", i+1)
		fmt.Printf("Rank: %d\n", ranking.Rank)
		fmt.Printf("Score: %f\n", ranking.Score)
		PrintRun(ranking.Run)
	}

}
