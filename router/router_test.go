package router

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"net/url"
	"production_service/models"
	"strings"
	"testing"
	"time"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetHome(t *testing.T) {
	// Mock a request sent, technically the request did not use HTTP layers
	// but when passed to the function, it mocks same behaviour
	r := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	w := httptest.NewRecorder()
	GetHome(w, r)
	assert.NotEmpty(t, w)

	response, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, response)
	assert.Equal(t, string(response), "\"I am Here! :)\"\n")
}

func TestGetUsers(t *testing.T) {
	// Mock a request sent, technically the request did not use HTTP layers
	// but when passed to the function, it mocks same behaviour
	r := httptest.NewRequest("GET", "http://localhost:8080/users", nil)
	w := httptest.NewRecorder()
	GetUsers(w, r)
	assert.NotEmpty(t, w)

	responseBody, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, responseBody)

	var target models.UserResponse
	json.Unmarshal([]byte(responseBody), &target)
	assert.NotEmpty(t, target)
	assert.Equal(t, target.Type, "success")
	assert.IsType(t, target.Data, []models.User{})
}

func TestCreateAndGetUser(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Mock a request sent, technically the request did not use HTTP layers
	// but when passed to the function, it mocks same behaviour
	data := url.Values{}
	data.Set("userid", fmt.Sprint(rand.Int()))
	data.Set("username", "foo")
	data.Set("userrole", "bar")
	r := httptest.NewRequest("POST", "http://localhost:8080/users/", strings.NewReader(data.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	CreateUser(w, r)
	assert.NotEmpty(t, w)

	responseBody, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, responseBody)

	var target models.UserResponse
	json.Unmarshal([]byte(responseBody), &target)
	assert.NotEmpty(t, target)
	assert.Equal(t, target.Type, "success")
	assert.IsType(t, target.Data, []models.User{})

	r = httptest.NewRequest("GET", "http://localhost:8080/users/foo", nil)
	r = mux.SetURLVars(r, map[string]string{"username": "foo"})
	w = httptest.NewRecorder()
	GetUser(w, r)

	responseBody, _ = ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, responseBody)

	json.Unmarshal([]byte(responseBody), &target)
	assert.NotEmpty(t, target)
	assert.Equal(t, target.Type, "success")
	assert.IsType(t, target.Data, []models.User{})
}

func TestGetHubs(t *testing.T) {
	// Mock a request sent, technically the request did not use HTTP layers
	// but when passed to the function, it mocks same behaviour
	r := httptest.NewRequest("GET", "http://localhost:8080/hubs", nil)
	w := httptest.NewRecorder()
	GetHubs(w, r)
	assert.NotEmpty(t, w)

	responseBody, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, responseBody)

	var target models.HubResponse
	json.Unmarshal([]byte(responseBody), &target)
	assert.NotEmpty(t, target)
	assert.Equal(t, target.Type, "success")
	assert.IsType(t, target.Data, []models.Hub{})
}

func TestCreateAndGetHub(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Mock a request sent, technically the request did not use HTTP layers
	// but when passed to the function, it mocks same behaviour
	data := url.Values{}
	data.Set("hubid", fmt.Sprint(rand.Int()))
	data.Set("hubname", "foo")
	data.Set("geolocation", "bar")
	r := httptest.NewRequest("POST", "http://localhost:8080/hubs", strings.NewReader(data.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	CreateHub(w, r)
	assert.NotEmpty(t, w)

	responseBody, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, responseBody)

	var target models.HubResponse
	json.Unmarshal([]byte(responseBody), &target)
	assert.NotEmpty(t, target)
	assert.Equal(t, target.Type, "success")
	assert.IsType(t, target.Data, []models.Hub{})

	r = httptest.NewRequest("GET", "http://localhost:8080/hubs/foo", nil)
	r = mux.SetURLVars(r, map[string]string{"hubname": "foo"})
	w = httptest.NewRecorder()
	GetHub(w, r)

	responseBody, _ = ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, responseBody)

	json.Unmarshal([]byte(responseBody), &target)
	assert.NotEmpty(t, target)
	assert.Equal(t, target.Type, "success")
	assert.IsType(t, target.Data, []models.Hub{})
}

func TestGetTeams(t *testing.T) {
	// Mock a request sent, technically the request did not use HTTP layers
	// but when passed to the function, it mocks same behaviour
	r := httptest.NewRequest("GET", "http://localhost:8080/teams", nil)
	w := httptest.NewRecorder()
	GetTeams(w, r)
	assert.NotEmpty(t, w)

	responseBody, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, responseBody)

	var target models.TeamResponse
	json.Unmarshal([]byte(responseBody), &target)
	assert.NotEmpty(t, target)
	assert.Equal(t, target.Type, "success")
	assert.IsType(t, target.Data, []models.Team{})
}

func TestCreateAndGetTeam(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Mock a request sent, technically the request did not use HTTP layers
	// but when passed to the function, it mocks same behaviour
	data := url.Values{}
	data.Set("teamid", fmt.Sprint(rand.Int()))
	data.Set("teamname", "foo")
	data.Set("teamtype", "bar")
	r := httptest.NewRequest("POST", "http://localhost:8080/teams", strings.NewReader(data.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	CreateTeam(w, r)
	assert.NotEmpty(t, w)

	responseBody, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, responseBody)

	var target models.TeamResponse
	json.Unmarshal([]byte(responseBody), &target)
	assert.NotEmpty(t, target)
	assert.Equal(t, target.Type, "success")
	assert.IsType(t, target.Data, []models.Team{})

	r = httptest.NewRequest("GET", "http://localhost:8080/teams/foo", nil)
	r = mux.SetURLVars(r, map[string]string{"hubname": "foo"})
	w = httptest.NewRecorder()
	GetHub(w, r)

	responseBody, _ = ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, responseBody)

	json.Unmarshal([]byte(responseBody), &target)
	assert.NotEmpty(t, target)
	assert.Equal(t, target.Type, "success")
	assert.IsType(t, target.Data, []models.Team{})
}
