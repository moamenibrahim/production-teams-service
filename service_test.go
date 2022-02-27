package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
)

func newTestServer(path string, h func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
    mux := http.NewServeMux()
    server := httptest.NewServer(mux)
    mux.HandleFunc(path, h)
    return server
}

type MySuite struct {
	suite.Suite
 }

 func (m *MySuite) SetupSuite() {
	log.Println("setup suite")
	go SetupHandlers()
 }
 
 func (m *MySuite) TearDownSuite() {
	log.Println("teardown suite")
 }

func (suite *MySuite) TestHomeEndpoint(){
	uri := "http://localhost:8080/"
	server := newTestServer(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	request, err := http.NewRequest("GET", uri, nil)
	assert.Equal(suite.T(), err, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.Equal(suite.T(), err, nil)
	assert.Equal(suite.T(), response.StatusCode, http.StatusOK)
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(suite.T(), string(responseBody), "\"I am Here! :)\"\n")

	client.CloseIdleConnections()
	defer response.Body.Close()
}


func (suite *MySuite) TestUsersEndpoint(){
	uri := "http://localhost:8080/users"
	server := newTestServer(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	request, err := http.NewRequest("GET", uri, nil)
	assert.Equal(suite.T(), err, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.Equal(suite.T(), err, nil)
	assert.Equal(suite.T(), response.StatusCode, http.StatusOK)

	client.CloseIdleConnections()
	defer response.Body.Close()
}


func (suite *MySuite) TestTeamsEndpoint(){
	uri := "http://localhost:8080/teams"
	server := newTestServer(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	request, err := http.NewRequest("GET", uri, nil)
	assert.Equal(suite.T(), err, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.Equal(suite.T(), err, nil)
	assert.Equal(suite.T(), response.StatusCode, http.StatusOK)

	client.CloseIdleConnections()
	defer response.Body.Close()
}


func (suite *MySuite) TestHubsEndpoint(){
	uri := "http://localhost:8080/hubs"
	server := newTestServer(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	request, err := http.NewRequest("GET", uri, nil)
	assert.Equal(suite.T(), err, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.Equal(suite.T(), err, nil)
	assert.Equal(suite.T(), response.StatusCode, http.StatusOK)

	client.CloseIdleConnections()
	defer response.Body.Close()
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(MySuite))
  }
