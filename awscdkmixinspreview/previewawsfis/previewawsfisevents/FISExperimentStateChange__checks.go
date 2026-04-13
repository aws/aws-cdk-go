//go:build !no_runtime_type_checking

package previewawsfisevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateFISExperimentStateChange_EventPatternParameters(options *FISExperimentStateChange_FISExperimentStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

