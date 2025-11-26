//go:build no_runtime_type_checking

package previewawscloudwatchevents

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlarmEvents) validateCloudWatchAlarmConfigurationChangePatternParameters(options *AlarmEvents_CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangeProps) error {
	return nil
}

func (a *jsiiProxy_AlarmEvents) validateCloudWatchAlarmStateChangePatternParameters(options *AlarmEvents_CloudWatchAlarmStateChange_CloudWatchAlarmStateChangeProps) error {
	return nil
}

func validateAlarmEvents_FromAlarmParameters(alarmRef interfacesawscloudwatch.IAlarmRef) error {
	return nil
}

