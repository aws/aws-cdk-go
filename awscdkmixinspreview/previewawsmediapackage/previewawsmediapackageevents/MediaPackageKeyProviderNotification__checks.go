//go:build !no_runtime_type_checking

package previewawsmediapackageevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateMediaPackageKeyProviderNotification_MediaPackageKeyProviderNotificationPatternParameters(options *MediaPackageKeyProviderNotification_MediaPackageKeyProviderNotificationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

