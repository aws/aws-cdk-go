//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func validateAlarmRule_AtLeastParameters(alarmState AlarmState, options *AtLeastOptions) error {
	return nil
}

func validateAlarmRule_AtLeastNotParameters(alarmState AlarmState, options *AtLeastOptions) error {
	return nil
}

func validateAlarmRule_FromAlarmParameters(alarm interfacesawscloudwatch.IAlarmRef, alarmState AlarmState) error {
	return nil
}

func validateAlarmRule_FromBooleanParameters(value *bool) error {
	return nil
}

func validateAlarmRule_FromStringParameters(alarmRule *string) error {
	return nil
}

func validateAlarmRule_NotParameters(operand IAlarmRule) error {
	return nil
}

