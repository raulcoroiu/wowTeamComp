package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/raulcoroiu/wowTeamComp/pkg/models"
)

const (
	apiURL = "https://raider.io/api/v1/mythic-plus/runs?season=season-df-1&region=world&page=0"
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

func PrintRuns(apiResponse *models.ApiResponse) {
	fmt.Println("Top runs for your spec:")

	for _, ranking := range apiResponse.Rankings {

		fmt.Println(SpecAndClassExist(ranking, "Druid", "Balance"))

	}

}

func SpecAndClassExist(vector models.Ranking, class string, spec string) Result {

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
		}
	}
	return results
}
