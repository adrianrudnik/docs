package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ory/client-go"
)

var ory *client.APIClient

func init() {
	cfg := client.NewConfiguration()
	cfg.Servers = client.ServerConfigurations{
		{URL: fmt.Sprintf("https://%s.projects.oryapis.com", os.Getenv("ORY_PROJECT_SLUG"))},
	}

	ory = client.NewAPIClient(cfg)
}

func GetLogin(ctx context.Context, flowId string) (*client.LoginFlow, error) {
	// highlight-start
	flow, _, err := ory.FrontendApi.GetLoginFlow(ctx).Id(flowId).Execute()
	if err != nil {
		return nil, err
	}

	return flow, nil
	// highlight-end
}
