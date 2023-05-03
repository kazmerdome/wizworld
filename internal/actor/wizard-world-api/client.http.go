package wizardworldapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ListElixirs     = "/Elixirs"
	ListIngredients = "/Ingredients"
)

type httpClient struct {
	url        string
	HttpClient http.Client
}

func NewHttpClient(timeoutInSeconds int, url string) *httpClient {
	client := http.Client{Timeout: time.Duration(timeoutInSeconds) * time.Second}
	return &httpClient{
		url:        url,
		HttpClient: client,
	}
}

func (r *httpClient) ListElixirs(ctx context.Context, params ListElixirsRequest) ([]ElixirResponse, error) {
	url := fmt.Sprintf("%s%s", r.url, ListElixirs)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	query := req.URL.Query()
	if params.Ingredient != "" {
		query.Add("Ingredient", params.Ingredient)
	}
	req.URL.RawQuery = query.Encode()

	response, err := r.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to GET data from wizard world api: communication error")
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(fmt.Sprintf("Unable to GET data from wizard world api, %v status received", response.StatusCode))
	}

	var er []ElixirResponse
	err = json.Unmarshal(responseBody, &er)
	if err != nil {
		return nil, err
	}

	return er, nil
}

func (r *httpClient) ListIngredients(ctx context.Context) ([]IngredientResponse, error) {
	url := fmt.Sprintf("%s%s", r.url, ListIngredients)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	query := req.URL.Query()
	req.URL.RawQuery = query.Encode()

	response, err := r.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to GET data from wizard world api: communication error")
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(fmt.Sprintf("Unable to GET data from wizard world api, %v status received", response.StatusCode))
	}

	var er []IngredientResponse
	err = json.Unmarshal(responseBody, &er)
	if err != nil {
		return nil, err
	}

	return er, nil
}
