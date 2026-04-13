//go:build !no_runtime_type_checking

package previewawsmediapackageevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateMediaPackageInputNotification_EventPatternParameters(options *MediaPackageInputNotification_MediaPackageInputNotificationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

