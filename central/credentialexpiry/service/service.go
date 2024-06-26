package service

import (
	"context"
	"crypto/tls"

	imageIntegrationStore "github.com/stackrox/rox/central/imageintegration/datastore"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/clientconn"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/grpc"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/mtls"
)

var (
	log = logging.LoggerForModule()
)

// Service provides the interface to the microservice that serves alert data.
type Service interface {
	grpc.APIService

	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)
	v1.CredentialExpiryServiceServer
}

// New returns a new Service instance using the given DataStore.
func New(imageIntegrations imageIntegrationStore.DataStore) Service {
	subjects := []mtls.Subject{mtls.ScannerSubject}
	if features.ScannerV4.Enabled() {
		subjects = append(subjects, mtls.ScannerV4IndexerSubject, mtls.ScannerV4MatcherSubject)
	}

	tlsConfigs := make(map[mtls.Subject]*tls.Config, len(subjects))
	for _, subject := range subjects {
		tlsConfig, err := clientconn.TLSConfig(subject, clientconn.TLSConfigOptions{
			UseClientCert: clientconn.MustUseClientCert,
		})
		if err != nil {
			// This case is hit if the Central CA cert cannot be loaded. This case is hit during some upgrade-tests
			// because in ancient versions, Central used to issue itself a cert on startup.
			// However, Central uses the exact same function to talk to scanner/Scanner V4, so any customer who actually uses
			// scanner/Scanner V4 must have patched their deployment to not hit this.
			// At the same time, we don't want to make this a fatal error, so just log a warning.
			log.Warnf("Failed to initialize %q TLS config: %v", subject.Identifier, err)
			continue
		}

		tlsConfigs[subject] = tlsConfig
	}

	return &serviceImpl{
		imageIntegrations: imageIntegrations,
		scannerConfigs:    tlsConfigs,
		expiryFunc:        maybeGetExpiryFromScannerAt,
	}
}
