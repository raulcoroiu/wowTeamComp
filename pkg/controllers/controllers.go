package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
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

func GetBestTeamHandler(c *gin.Context) {
	// Parse query parameters for class and spec
	class := c.Query("class")
	spec := c.Query("spec")

	if class == "" || spec == "" {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "Missing class or spec query parameter"})
		c.HTML(http.StatusBadRequest, "index.html", gin.H{"error": "Missing class or spec"})
		return
	}

	// Make the request to the Raider.io API (you may want to add error handling)
	responseBody, _ := MakeRequest()
	apiResponse, _ := ParseResponse(responseBody)

	// Find the best team for the specified class and spec
	var bestTeam Result
	for _, ranking := range apiResponse.Rankings {
		bestTeam = SpecAndClassExist(ranking, class, spec)
		if bestTeam.Rank > 0 {
			break
		}
	}

	if bestTeam.Rank == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No team found for the specified class and spec"})
	} else {
		c.JSON(http.StatusOK, bestTeam)
	}
}
