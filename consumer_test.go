package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestConsumer(t *testing.T) {
	type User struct {
		Firstname   string `json:"firstname" pact:"example=Bill"`
		Lastname    string `json:"lastname" pact:"example=Muray"`
		Age         int    `json:"age" pact:"example=68"`
		Catchphrase string `json:"catchphrase" pact:"example=Ghostbuster..."`
	}

	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "MyConsumer",
		Provider: "MyProvider",
		Host:     "localhost",
	}
	defer pact.Teardown()

	// Pass in test case. This is the component that makes the external HTTP call
	var test = func() (err error) {
		u := fmt.Sprintf("http://localhost:%d/", pact.Server.Port)
		req, err := http.NewRequest("GET", u, strings.NewReader(""))
		if err != nil {
			return
		}

		// NOTE: by default, request bodies are expected to be sent with a Content-Type
		// of application/json. If you don't explicitly set the content-type, you
		// will get a mismatch during Verification.
		req.Header.Set("Content-Type", "application/json")

		if _, err = http.DefaultClient.Do(req); err != nil {
			return
		}

		return
	}

	// Set up our expected interactions.
	pact.
		AddInteraction().
		Given("Root user exists").
		UponReceiving("A request to get root user").
		WithRequest(dsl.Request{
			Method:  "GET",
			Path:    dsl.String("/"),
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Match(&User{}),
		})

	// Run the test, verify it did what we expected and capture the contract
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}

	return
}
