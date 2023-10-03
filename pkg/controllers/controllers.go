package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulcoroiu/wowTeamComp/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/internal/errors"
)

const (
	apiURL = "https://raider.io/api/v1/mythic-plus/runs?season=season-df-1&region=world&page=0"
)

type UserService struct {
}

type authUser struct {
	email        string
	passwordHash string
}

type User struct {
	Email    string
	Password string
}

type Member struct {
	Class string `json:"class"`
	Spec  string `json:"spec"`
}

type Result struct {
	Rank    int64    `json:"rank"`
	Members []Member `json:"members"`
}

var authUserDB = map[string]authUser{}
var DefaultUserService UserService

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

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json")

	class := c.Query("class")
	spec := c.Query("spec")

	if class == "" || spec == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing class or spec query parameter"})
		return
	}

	responseBody, _ := MakeRequest()
	apiResponse, _ := ParseResponse(responseBody)

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
		bestTeamJSON, err := json.Marshal(bestTeam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		}
		c.Data(http.StatusOK, "application/json", bestTeamJSON)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sing-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w, r)
	case "/sign-up-form":
		getSignUpPage(w, r)
	}
}

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	_, ok := DefaultUserService.VerifyUser(newUser)
	if !ok {
		fileName := "sign-up.html"
		t, _ := template.ParseFiles(filename)
		t.ExecuteTemplate(w, fileName, "User Sign-in Failure")
		return
	}

	filename := "sign-up.html"
	t, _ := template.ParseFiles(filename)
	t.ExecuteTemplate(w, fileName, "User Sign-in")

}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := DefaultUserService.createUser(newUser)
	if err != nil {
		fileName := "sign-up.html"
		t, _ := template.ParseFiles(filename)
		t.ExecuteTemplate(w, fileName, "New User Sign-up Failure")
		return
	}

	filename := "sign-up.html"
	t, _ := template.ParseFiles(filename)
	t.ExecuteTemplate(w, fileName, "New User Sign-up")

}

func getUser(r *http.Request) User {
	email := r.FormValue("email")
	password := r.FormValue("password")

	return User{
		Email:    email,
		Password: password,
	}
}

func (UserService) createUser(newUser User) error {
	_, ok := authUserDB[newUser.Email]
	if !ok {
		return errors.New("user already exist")
	}

	passwordHash, err := getPasswordHash(newUser.Password)
	if err != nil {
		return err
	}
	newAuthUser := authUser{
		email:        newUser.Email,
		passwordHash: passwordHash,
	}

	authUserDB[newAuthUser.email] = newAuthUser

	return nil

}

func getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err
}

func (UserService) VerifyUser(user User) bool{
	authUser, ok := authUserDB[user.Email]
	if !ok {
		return false
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(authUser.passwordHash),
		[]byte(user.Password))

	return err == nil



}