//go:build !no_runtime_type_checking

package previewawsiotanalyticsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateIoTAnalyticsDataSetLifeCycleNotification_EventPatternParameters(options *IoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

