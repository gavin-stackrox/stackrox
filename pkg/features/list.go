package features

//lint:file-ignore U1000 we want to introduce this feature flag unused.

var (
	// AdmissionControlService enables running admission control as a separate microservice.
	AdmissionControlService = registerFeature("Separate admission control microservice", "ROX_ADMISSION_CONTROL_SERVICE", true)

	// csvExport enables CSV export of search results.
	csvExport = registerFeature("Enable CSV export of search results", "ROX_CSV_EXPORT", false)

	// SupportSlimCollectorMode enables support for retrieving slim Collector bundles from central.
	SupportSlimCollectorMode = registerFeature("Support slim Collector mode", "ROX_SUPPORT_SLIM_COLLECTOR_MODE", true)

	// AwsSecurityHubIntegration enables the AWS Security Hub Integration UI.
	AwsSecurityHubIntegration = registerFeature("Show AWS Security Hub Integration in UI", "ROX_AWS_SECURITY_HUB_INTEGRATION", true)

	// NetworkGraphPorts enables port-related features in the network graph.
	NetworkGraphPorts = registerFeature("Enable port-related features in network graph", "ROX_NETWORK_GRAPH_PORTS", true)

	// NetworkFlowsSearchFilterUI enables client-side filtering for network flows.
	// NB: When removing this feature flag, remove references in ui/src/utils/featureFlags.js
	NetworkFlowsSearchFilterUI = registerFeature("Enable client-side network flows search", "ROX_NETWORK_FLOWS_SEARCH_FILTER_UI", true)

	// ComplianceInRocksDB switches compliance over to using RocksDB instead of Bolt
	ComplianceInRocksDB = registerFeature("Switch compliance to using RocksDB", "ROX_COMPLIANCE_IN_ROCKSDB", true)

	// SyslogIntegration enables UI for a Syslog integration in the Integrations section.
	SyslogIntegration = registerFeature("Enable UI for a Syslog integration in the Integrations section", "ROX_SYSLOG_INTEGRATION", true)

	// SensorInstallationExperience enables new features related to the new installation experience for sensor.
	SensorInstallationExperience = registerFeature("Enable new installation user experience for Sensor", "ROX_SENSOR_INSTALLATION_EXPERIENCE", true)

	// NetworkDetection enables new features related to the new network detection experience.
	NetworkDetection = registerFeature("Enable new network detection experience", "ROX_NETWORK_DETECTION", true)

	// NetworkDetectionBaselineViolation enables new features related to the baseline violation part of the network detection experience.
	NetworkDetectionBaselineViolation = registerFeature("Enable network detection baseline violation", "ROX_NETWORK_DETECTION_BASELINE_VIOLATION", false)

	// HostScanning enables new features related to the new host scanning experience in VM.
	HostScanning = registerFeature("Enable new host scanning experience", "ROX_HOST_SCANNING", false)

	// SensorTLSChallenge enables Sensor to receive Centrals configured additional-ca an default certs.
	SensorTLSChallenge = registerFeature("Enable Sensor to receive default and additional CA certificates from Central", "ROX_SENSOR_TLS_CHALLENGE", true)

	// K8sEventDetection enables detection of kubernetes events.
	K8sEventDetection = registerFeature("Enable detection of kubernetes events", "ROX_K8S_EVENTS_DETECTION", true)

	// IntegrationsAsConfig enables loading integrations from config
	IntegrationsAsConfig = registerFeature("Enable loading integrations from config", "ROX_INTEGRATIONS_AS_CONFIG", false)

	// ScopedAccessControl enables scoped access control in core product
	ScopedAccessControl = registerFeature("Enable scoped access control in core product", "ROX_SCOPED_ACCESS_CONTROL_V2", false)
)
