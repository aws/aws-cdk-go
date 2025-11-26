//go:build !no_runtime_type_checking

package previewawscloudwatchevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
)

func (a *jsiiProxy_AlarmEvents) validateCloudWatchAlarmConfigurationChangePatternParameters(options *AlarmEvents_CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AlarmEvents) validateCloudWatchAlarmStateChangePatternParameters(options *AlarmEvents_CloudWatchAlarmStateChange_CloudWatchAlarmStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAlarmEvents_FromAlarmParameters(alarmRef interfacesawscloudwatch.IAlarmRef) error {
	if alarmRef == nil {
		return fmt.Errorf("parameter alarmRef is required, but nil was provided")
	}

	return nil
}

