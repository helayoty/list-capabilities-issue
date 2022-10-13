package client

import (
	"context"

	azaci "github.com/Azure/azure-sdk-for-go/services/containerinstance/mgmt/2021-10-01/containerinstance"
	"github.com/pkg/errors"
)

func (config *Config) ListCapabilities(ctx context.Context, region string) (*[]azaci.Capabilities, error) {
	capabilities, err := config.lClient.ListCapabilitiesComplete(ctx, region)

	if err != nil {
		return nil, errors.Wrapf(err, "unable to fetch the ACI capabilities for the location %s", region)
	}

	result := capabilities.Response().Value
	if result == nil {
		return nil, errors.Wrapf(err, " result is empty")
	}
	return result, nil
}
