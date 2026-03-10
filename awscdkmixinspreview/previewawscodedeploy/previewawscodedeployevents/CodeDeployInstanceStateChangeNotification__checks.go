//go:build !no_runtime_type_checking

package previewawscodedeployevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCodeDeployInstanceStateChangeNotification_CodeDeployInstanceStateChangeNotificationPatternParameters(options *CodeDeployInstanceStateChangeNotification_CodeDeployInstanceStateChangeNotificationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

