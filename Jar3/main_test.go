// main_test.go
package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchScimResponse(t *testing.T) {
	// Start HTTP mocking
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the GET request to the API endpoint
	httpmock.RegisterResponder("GET", "https://api.example.com/scim/v2/Deployments",
		httpmock.NewStringResponder(200, `{
			"Resources": [
				{"DeploymentInfo": {"Region": "us-east-1"}},
				{"DeploymentInfo": {"Region": "eu-central-1"}},
				{"DeploymentInfo": {"Region": "ap-northeast-1"}}
			]
		}`))

	// Call the function under test
	scimResponse, err := FetchScimResponse("https://api.example.com/scim/v2/Deployments")

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the response is as expected
	expectedRegions := []string{"us-east-1", "eu-central-1", "ap-northeast-1"}
	for i, region := range expectedRegions {
		assert.Equal(t, region, scimResponse.Resources[i].DeploymentInfo.Region)
	}

	// Assert that the correct URL was called
	assert.Equal(t, 1, httpmock.GetTotalCallCount(), "The API endpoint was not called exactly once")
}
