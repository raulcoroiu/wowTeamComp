package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/raulcoroiu/wowTeamComp/pkg/models"
)

const (
	apiURL = "https://raider.io/api/v1/mythic-plus/runs?season=season-df-1&region=world&dungeon=all&page=0"
)

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

//---De aci am incercat cu  chat GPT sa fac ceva :)))

func getCompositionFromRoster(roster []models.Roster) []models.ClassFaction {
	composition := make([]models.ClassFaction, 0)

	for _, r := range roster {
		composition = append(composition, r.Character.Faction)
	}

	return composition
}

func FindTopCompositions(apiResponse *models.ApiResponse, classSpec string) map[int64][]models.CompositionData {
	compositionsByKeystone := make(map[int64][]models.CompositionData)

	for _, ranking := range apiResponse.Rankings {
		run := ranking.Run

		for _, roster := range run.Roster {
			if roster.Character.Spec.Name == classSpec {
				keystoneLevel := run.MythicLevel
				composition := getCompositionFromRoster(run.Roster)
				averageScore := ranking.Score

				compositionsByKeystone[keystoneLevel] = append(compositionsByKeystone[keystoneLevel], models.CompositionData{
					Composition:  composition,
					AverageScore: averageScore,
				})

				sort.Slice(compositionsByKeystone[keystoneLevel], func(i, j int) bool {
					return compositionsByKeystone[keystoneLevel][i].AverageScore > compositionsByKeystone[keystoneLevel][j].AverageScore
				})

				if len(compositionsByKeystone[keystoneLevel]) > 3 {
					compositionsByKeystone[keystoneLevel] = compositionsByKeystone[keystoneLevel][:3]
				}
			}
		}
	}

	return compositionsByKeystone
}
