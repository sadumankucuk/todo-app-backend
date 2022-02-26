package pact

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
	"todo/server"
)

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	svr := server.NewServer()
	go svr.StartServer(port)

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "TodoService",
		Consumer:                 "TodoApp",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs: []string{
			"/Users/sadumankucuk/Desktop/frontend/pact/pacts/todoapp-todoservice.json",
		},
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponses), "pact tests run")
}
