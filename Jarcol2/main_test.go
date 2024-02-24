// main_test.go
package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchMessage(t *testing.T) {
	// Start HTTP mocking
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the GET request to the hardcoded URL
	httpmock.RegisterResponder("GET", apiURL,
		httpmock.NewStringResponder(200, `{"message": "Hello, World!"}`))

	// Call the function
	message, err := FetchMessage()

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert the fetched message is as expected
	assert.Equal(t, "Hello, World!", message)

	// Assert that the correct URL was called
	info := httpmock.GetCallCountInfo()
	assert.Equal(t, 1, info["GET "+apiURL], "The API endpoint was not called exactly once")
}
