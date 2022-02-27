package main

import (
	"log"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

func newTestServer(path string, h func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
    mux := http.NewServeMux()
    server := httptest.NewServer(mux)
    mux.HandleFunc(path, h)
    return server
}

// You can use testing.T, if you want to test the code without benchmarking
func setupSuite(tb testing.TB) func(tb testing.TB) {
	log.Println("setup suite")
	go SetupHandlers()

	// Return a function to teardown the test
	return func(tb testing.TB) {
		log.Println("teardown suite")
	}
}

func TestHomeEndpoint(t *testing.T){
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	uri := "http://localhost:8080/"
	server := newTestServer(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	request, err := http.NewRequest("GET", uri, nil)
	assert.Equal(t, err, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.Equal(t, err, nil)
	assert.Equal(t, response.StatusCode, http.StatusOK)
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, string(responseBody), "\"I am Here! :)\"\n")

	client.CloseIdleConnections()
	defer response.Body.Close()
}


func TestUsersEndpoint(t *testing.T){
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	uri := "http://localhost:8080/users"
	server := newTestServer(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	request, err := http.NewRequest("GET", uri, nil)
	assert.Equal(t, err, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.Equal(t, err, nil)
	assert.Equal(t, response.StatusCode, http.StatusOK)

	client.CloseIdleConnections()
	defer response.Body.Close()
}


func TestTeamsEndpoint(t *testing.T){
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	uri := "http://localhost:8080/teams"
	server := newTestServer(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	request, err := http.NewRequest("GET", uri, nil)
	assert.Equal(t, err, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.Equal(t, err, nil)
	assert.Equal(t, response.StatusCode, http.StatusOK)

	client.CloseIdleConnections()
	defer response.Body.Close()
}


func TestHubsEndpoint(t *testing.T){
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	uri := "http://localhost:8080/hubs"
	server := newTestServer(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	request, err := http.NewRequest("GET", uri, nil)
	assert.Equal(t, err, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.Equal(t, err, nil)
	assert.Equal(t, response.StatusCode, http.StatusOK)

	client.CloseIdleConnections()
	defer response.Body.Close()
}
