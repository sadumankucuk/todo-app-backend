package pact

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"strconv"
	"testing"
	"todo/server"
)

func TestProvider(t *testing.T) {
	//t.Skip("skipping testing in short mode")
	port, _ := utils.GetFreePort()
	svr := server.NewServer()
	go svr.StartServer(strconv.Itoa(port))

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "TodoService",
		Consumer:                 "TodoApp",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs: []string{
			"https://bootcamp.pactflow.io/pacts/provider/TodoService/consumer/TodoApp/latest",
		},
		BrokerToken:                "LYp12gGXhTdQ4TadedFYdw",
		BrokerURL:                  "https://bootcamp.pactflow.io",
		ProviderVersion:            "v1",
		PublishVerificationResults: true,
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponses), "pact tests run")
}
