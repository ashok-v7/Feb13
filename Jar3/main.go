package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DeploymentInfo struct {
	Region string `json:"Region"`
}

type ScimResponse struct {
	Resources []struct {
		DeploymentInfo DeploymentInfo `json:"DeploymentInfo"`
	} `json:"Resources"`
}

// FetchScimResponse makes an HTTP GET request and returns a ScimResponse.
func FetchScimResponse(apiURL string) (*ScimResponse, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var scimResponse ScimResponse
	err = json.Unmarshal(bodyBytes, &scimResponse)
	if err != nil {
		return nil, err
	}

	return &scimResponse, nil
}

func main() {
	apiURL := "https://api.example.com/scim/v2/Deployments"

	scimResponse, err := FetchScimResponse(apiURL)
	if err != nil {
		log.Fatalf("Error fetching SCIM response: %v", err)
	}

	fmt.Println("SCIM Deployment Info:")
	for _, resource := range scimResponse.Resources {
		fmt.Printf("Region: %s\n", resource.DeploymentInfo.Region)
	}
}
