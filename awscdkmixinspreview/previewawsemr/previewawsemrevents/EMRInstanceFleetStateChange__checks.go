//go:build !no_runtime_type_checking

package previewawsemrevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEMRInstanceFleetStateChange_EventPatternParameters(options *EMRInstanceFleetStateChange_EMRInstanceFleetStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

