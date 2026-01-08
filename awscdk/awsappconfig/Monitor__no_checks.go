//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func validateMonitor_FromCfnMonitorsPropertyParameters(monitorsProperty *CfnEnvironment_MonitorsProperty) error {
	return nil
}

func validateMonitor_FromCloudWatchAlarmParameters(alarm interfacesawscloudwatch.IAlarmRef) error {
	return nil
}

