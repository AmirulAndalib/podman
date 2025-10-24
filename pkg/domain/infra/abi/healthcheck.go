//go:build !remote

package abi

import (
	"context"

	"github.com/containers/podman/v6/libpod/define"
	"github.com/containers/podman/v6/pkg/domain/entities"
)

func (ic *ContainerEngine) HealthCheckRun(ctx context.Context, nameOrID string, _ entities.HealthCheckOptions) (*define.HealthCheckResults, error) {
	status, err := ic.Libpod.HealthCheck(ctx, nameOrID)
	if err != nil {
		return nil, err
	}
	report := define.HealthCheckResults{
		Status: status.String(),
	}
	return &report, nil
}
