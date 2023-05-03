package wizardworldapi_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	wizardworldapi "github.com/kazmerdome/wizworld/internal/actor/wizard-world-api"
	"github.com/stretchr/testify/assert"
)

type clientFixture struct {
	mockServer  *httptest.Server
	elixirs     []wizardworldapi.ElixirResponse
	ingredients []wizardworldapi.IngredientResponse
}

func newClientFixture(t *testing.T) *clientFixture {
	f := new(clientFixture)
	f.ingredients = []wizardworldapi.IngredientResponse{
		{
			Id:   "1",
			Name: "Test Ingredient 1",
		},
		{
			Id:   "2",
			Name: "Test Ingredient 2",
		},
		{
			Id:   "3",
			Name: "Test Ingredient 3",
		},
	}
	f.elixirs = []wizardworldapi.ElixirResponse{
		{
			Id:         "1",
			Name:       "Test Elixir 1",
			Effect:     "Test effect 1",
			Difficulty: "Easy",
			Ingredients: []wizardworldapi.IngredientResponse{
				{
					Id:   "1",
					Name: "Test Ingredient 1",
				},
				{
					Id:   "2",
					Name: "Test Ingredient 2",
				},
			},
		},
		{
			Id:         "2",
			Name:       "Test Elixir 2",
			Effect:     "Test effect 2",
			Difficulty: "Medium",
			Ingredients: []wizardworldapi.IngredientResponse{
				{
					Id:   "3",
					Name: "Test Ingredient 3",
				},
				{
					Id:   "4",
					Name: "Test Ingredient 4",
				},
			},
		},
	}
	f.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Elixirs" {
			jsonData, _ := json.Marshal(f.elixirs)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		} else if r.URL.Path == "/Ingredients" {
			jsonData, _ := json.Marshal(f.ingredients)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}))

	return f
}

// ListElixirs
//

func TestListElixirs_FailsOn_Not200(t *testing.T) {
	// Create a mock server for the API
	f := newClientFixture(t)
	f.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Elixirs" {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer f.mockServer.Close()

	// Create a new instance of the httpClient using the mock server URL
	httpClient := wizardworldapi.NewHttpClient(5, f.mockServer.URL)

	// Call the ListElixirs method with a test ingredient
	elixirs, err := httpClient.ListElixirs(context.Background(), wizardworldapi.ListElixirsRequest{})
	assert.EqualError(t, err, "Unable to GET data from wizard world api, 500 status received")
	assert.Empty(t, elixirs)
}

func TestListElixirs_FailsOn_Unmarshal(t *testing.T) {
	// Create a mock server for the API
	f := newClientFixture(t)
	f.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Elixirs" {
			test := struct {
				Test string
				X    int
			}{Test: "test", X: 3}
			jsonData, err := json.Marshal(test)
			assert.NoError(t, err)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}))
	defer f.mockServer.Close()

	// Create a new instance of the httpClient using the mock server URL
	httpClient := wizardworldapi.NewHttpClient(5, f.mockServer.URL)

	// Call the ListElixirs method with a test ingredient
	elixirs, err := httpClient.ListElixirs(context.Background(), wizardworldapi.ListElixirsRequest{})
	assert.EqualError(t, err, "json: cannot unmarshal object into Go value of type []wizardworldapi.ElixirResponse")
	assert.Empty(t, elixirs)
}

func TestListElixirs_FailsOn_HTTPError(t *testing.T) {
	// Create a mock server for the API
	f := newClientFixture(t)
	f.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Elixirs" {
			w.WriteHeader(http.StatusOK)
			jsonData, _ := json.Marshal(f.elixirs)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}))
	defer f.mockServer.Close()

	// Create a new instance of the httpClient using the mock server URL
	httpClient := wizardworldapi.NewHttpClient(5, f.mockServer.URL)

	// Override the httpClient with a client that will return an error
	httpClient.HttpClient = http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse("http://invalid-proxy")
			},
		},
	}
	// Call the ListElixirs method with a test ingredient
	elixirs, err := httpClient.ListElixirs(context.Background(), wizardworldapi.ListElixirsRequest{})
	assert.EqualError(t, err, "unable to GET data from wizard world api: communication error")
	assert.Empty(t, elixirs)
}

func TestListElixirs_Success(t *testing.T) {
	// Create a mock server for the API
	f := newClientFixture(t)
	defer f.mockServer.Close()

	// Create a new instance of the httpClient using the mock server URL
	httpClient := wizardworldapi.NewHttpClient(5, f.mockServer.URL)

	// Call the ListElixirs method with a test ingredient
	elixirs, err := httpClient.ListElixirs(context.Background(), wizardworldapi.ListElixirsRequest{Ingredient: "test-ingredient"})
	assert.NoError(t, err)
	assert.Equal(t, elixirs, f.elixirs)
}

// ListIngredients
//

func TestListIngredients_FailsOn_Not200(t *testing.T) {
	// Create a mock server for the API
	f := newClientFixture(t)
	f.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Ingredients" {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer f.mockServer.Close()

	// Create a new instance of the httpClient using the mock server URL
	httpClient := wizardworldapi.NewHttpClient(5, f.mockServer.URL)

	// Call the ListIngredients method
	ingredients, err := httpClient.ListIngredients(context.Background())
	assert.EqualError(t, err, "Unable to GET data from wizard world api, 500 status received")
	assert.Empty(t, ingredients)
}

func TestListIngredients_FailsOn_Unmarshal(t *testing.T) {
	// Create a mock server for the API
	f := newClientFixture(t)
	f.ingredients = []wizardworldapi.IngredientResponse{
		{
			Id:   "1",
			Name: "Test Ingredient 1",
		},
		{
			Id:   "2",
			Name: "Test Ingredient 2",
		},
	}
	f.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Ingredients" {
			test := struct {
				Test string
				X    int
			}{Test: "test", X: 3}
			jsonData, err := json.Marshal(test)
			assert.NoError(t, err)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}))
	defer f.mockServer.Close()

	// Create a new instance of the httpClient using the mock server URL
	httpClient := wizardworldapi.NewHttpClient(5, f.mockServer.URL)

	// Call the ListIngredients method
	ingredients, err := httpClient.ListIngredients(context.Background())
	assert.EqualError(t, err, "json: cannot unmarshal object into Go value of type []wizardworldapi.IngredientResponse")
	assert.Empty(t, ingredients)
}

func TestListIngredients_FailsOn_HTTPError(t *testing.T) {
	// Create a mock server for the API
	f := newClientFixture(t)
	f.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Ingredients" {
			w.WriteHeader(http.StatusOK)
			jsonData, _ := json.Marshal(f.ingredients)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}))
	defer f.mockServer.Close()

	// Create a new instance of the httpClient using the mock server URL
	httpClient := wizardworldapi.NewHttpClient(5, f.mockServer.URL)

	// Override the httpClient with a client that will return an error
	httpClient.HttpClient = http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse("http://invalid-proxy")
			},
		},
	}

	// Call the ListIngredients method
	ingredients, err := httpClient.ListIngredients(context.Background())
	assert.EqualError(t, err, "unable to GET data from wizard world api: communication error")
	assert.Empty(t, ingredients)
}

func TestListIngredients_Success(t *testing.T) {
	// Create a mock server for the API
	f := newClientFixture(t)
	defer f.mockServer.Close()

	// Create a new instance of the httpClient using the mock server URL
	httpClient := wizardworldapi.NewHttpClient(5, f.mockServer.URL)

	// Call the ListIngredients method
	ingredients, err := httpClient.ListIngredients(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, ingredients, f.ingredients)
}
