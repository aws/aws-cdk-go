//go:build !no_runtime_type_checking

package previewawsemrevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEMRConfigurationError_EventPatternParameters(options *EMRConfigurationError_EMRConfigurationErrorProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

