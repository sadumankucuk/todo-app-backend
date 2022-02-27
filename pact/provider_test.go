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
		BrokerToken:     "LYp12gGXhTdQ4TadedFYdw",
		BrokerURL:       "https://bootcamp.pactflow.io",
		ProviderVersion: "v1",
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponses), "pact tests run")
}
