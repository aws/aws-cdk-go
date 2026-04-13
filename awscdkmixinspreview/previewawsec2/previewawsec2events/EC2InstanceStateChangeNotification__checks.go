//go:build !no_runtime_type_checking

package previewawsec2events

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEC2InstanceStateChangeNotification_EventPatternParameters(options *EC2InstanceStateChangeNotification_EC2InstanceStateChangeNotificationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

