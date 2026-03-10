//go:build !no_runtime_type_checking

package previewawscloudwatchevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangePatternParameters(options *CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

