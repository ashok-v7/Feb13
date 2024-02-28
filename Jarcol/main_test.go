// // main_test.go
package main

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchMessage(t *testing.T) {
	// Start HTTP mocking
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the GET request
	httpmock.RegisterResponder("GET", "https://api.example.com",
		httpmock.NewStringResponder(200, `{"message": "Hello, World!"}`))

	// Call the function
	message, err := FetchMessage("https://api.example.com")

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert the fetched message is as expected
	assert.Equal(t, "Hello, World!", message)

	// Assert that the correct URL was called
	info := httpmock.GetCallCountInfo()
	assert.Equal(t, 1, info["GET https://api.example.com"], "The API endpoint was not called exactly once")
	fmt.Println("Mock server call count info:", info)
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestFetchMessage(t *testing.T) {
// 	expectedMessage := "test message"
// 	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
// 		rw.Write([]byte(`{"message":"` + expectedMessage + `"}`))
// 	}))
// 	fmt.Println("server", server)
// 	defer server.Close()

// 	fmt.Println("server.url", server.URL)
// 	message, err := FetchMessage(server.URL)
// 	if err != nil {
// 		t.Errorf("FetchMessage returned an error: %v", err)
// 	}
// 	fmt.Println("Fetched Message:", message)
// 	fmt.Println("Expected Message:", expectedMessage)
// 	if message != expectedMessage {
// 		t.Errorf("FetchMessage returned unexpected message, got: %s, want: %s", message, expectedMessage)
// 	}
// }
