//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
)

func validateAlarmRule_AtLeastParameters(alarmState AlarmState, options *AtLeastOptions) error {
	if alarmState == "" {
		return fmt.Errorf("parameter alarmState is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAlarmRule_AtLeastNotParameters(alarmState AlarmState, options *AtLeastOptions) error {
	if alarmState == "" {
		return fmt.Errorf("parameter alarmState is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAlarmRule_FromAlarmParameters(alarm interfacesawscloudwatch.IAlarmRef, alarmState AlarmState) error {
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

