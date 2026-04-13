//go:build !no_runtime_type_checking

package previewawskmsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateKMSCMKRotation_EventPatternParameters(options *KMSCMKRotation_KMSCMKRotationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

