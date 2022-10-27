//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"
)

func validateAlarmRule_FromAlarmParameters(alarm IAlarm, alarmState AlarmState) error {
	if alarm == nil {
		return fmt.Errorf("parameter alarm is required, but nil was provided")
	}

	if alarmState == "" {
		return fmt.Errorf("parameter alarmState is required, but nil was provided")
	}

	return nil
}

func validateAlarmRule_FromBooleanParameters(value *bool) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateAlarmRule_FromStringParameters(alarmRule *string) error {
	if alarmRule == nil {
		return fmt.Errorf("parameter alarmRule is required, but nil was provided")
	}

	return nil
}

func validateAlarmRule_NotParameters(operand IAlarmRule) error {
	if operand == nil {
		return fmt.Errorf("parameter operand is required, but nil was provided")
	}

	return nil
}

