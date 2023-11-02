//go:build no_runtime_type_checking

package awscdkappconfigalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateMonitor_FromCfnMonitorsPropertyParameters(monitorsProperty *awsappconfig.CfnEnvironment_MonitorsProperty) error {
	return nil
}

func validateMonitor_FromCloudWatchAlarmParameters(alarm awscloudwatch.IAlarm) error {
	return nil
}

