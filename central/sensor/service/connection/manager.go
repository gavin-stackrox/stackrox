package connection

import (
	"context"
	"time"

	"github.com/stackrox/rox/central/sensor/service/pipeline"
	"github.com/stackrox/rox/generated/internalapi/central"
)

// CheckInRecorder updates the cluster contact time
type CheckInRecorder interface {
	UpdateClusterContactTime(ctx context.Context, clusterID string, time time.Time) error
}

// Manager is responsible for managing all active connections from sensors.
//go:generate mockgen-wrapper Manager
type Manager interface {
	HandleConnection(ctx context.Context, clusterID string, pf pipeline.Factory, server central.SensorService_CommunicateServer, recorder CheckInRecorder) error
	GetConnection(clusterID string) SensorConnection

	GetActiveConnections() []SensorConnection
}
