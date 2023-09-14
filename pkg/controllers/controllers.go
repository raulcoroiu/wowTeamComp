package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/raulcoroiu/wowTeamComp/pkg/models"
)

const (
	apiURL = "https://raider.io/api/v1/mythic-plus/runs?season=season-df-1&region=world&page=0"
	//apiURL = "https://raider.io/api/v1/mythic-plus/runs?season=season-7.3.0&region=world&dungeon=all"
)

type Member struct {
	Class string `json:"class"`
	Spec  string `json:"spec"`
}

type Result struct {
	Rank    int64    `json:"rank"`
	Members []Member `json:"members"`
}

func MakeRequest() ([]byte, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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

//func PrintRun(run models.Run) {
//	fmt.Println("Run:")
//	fmt.Printf("Season: %s\n", run.Season)
//	fmt.Printf("Status: %s\n", run.Status)
//	fmt.Printf("Keystone Run ID: %d\n", run.KeystoneRunID)
//	//fmt.Printf("keystone Run Name: %s\n", run.key)
//	fmt.Printf("Mythic Level: %d\n", run.MythicLevel)
//	fmt.Printf("Clear Time (ms): %d\n", run.ClearTimeMS)
//	fmt.Printf("Keystone Time (ms): %d\n", run.KeystoneTimeMS)
//	fmt.Printf("Completed At: %s\n", run.CompletedAt)
//	fmt.Printf("Number of Chests: %d\n", run.NumChests)
//	fmt.Printf("Time Remaining (ms): %d\n", run.TimeRemainingMS)

//	for _, member := range run.Roster {
//		fmt.Printf("Race: %s\n", member.Character.Race.Name)
//		fmt.Printf("Class: %s\n", member.Character.Class.Name)
//		fmt.Printf("Spec: %s\n", member.Character.Spec.Name)
//	}
//
//	// mai trebuie adaugate
//}

func PrintRuns(apiResponse *models.ApiResponse) {
	fmt.Println("Mythic+ Runs:")

	for _, ranking := range apiResponse.Rankings {
		//fmt.Printf("Ranking %d:\n", i+1)
		//fmt.Printf("Rank: %d\n", ranking.Rank)
		//fmt.Printf("Score: %f\n", ranking.Score)
		//PrintRun(ranking.Run)
		SpecAndClassExist(ranking, "Druid", "Balance")

	}

}

func SpecAndClassExist(vector models.Ranking, class string, spec string) {

	for _, member := range vector.Run.Roster {
		if member.Character.Class.Name == class && member.Character.Spec.Name == spec {

			fmt.Println("Global keystone Rank is:" + strconv.FormatInt(vector.Rank, 10))
			for i := 0; i <= 4; i++ {
				fmt.Println("Class " + strconv.Itoa(i+1) + ": " + vector.Run.Roster[i].Character.Class.Name)
				fmt.Println("Spec " + strconv.Itoa(i+1) + ": " + vector.Run.Roster[i].Character.Spec.Name)
			}

		}

	}

}

func BestCompHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request and retrieve the necessary data
	var vector models.Ranking
	// Implement code to populate 'vector' from the request, e.g., by decoding JSON from the request body

	class := r.URL.Query().Get("class")
	spec := r.URL.Query().Get("spec")

	var results Result

	for _, member := range vector.Run.Roster {
		if member.Character.Class.Name == class && member.Character.Spec.Name == spec {
			results.Rank = vector.Rank

			for i := 0; i <= 4; i++ {
				member := Member{
					Class: vector.Run.Roster[i].Character.Class.Name,
					Spec:  vector.Run.Roster[i].Character.Spec.Name,
				}
				results.Members = append(results.Members, member)
			}

			break
		}
	}

	// Convert the result to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetComp() {

}
