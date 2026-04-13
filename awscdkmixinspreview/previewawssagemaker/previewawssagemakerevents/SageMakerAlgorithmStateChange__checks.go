//go:build !no_runtime_type_checking

package previewawssagemakerevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateSageMakerAlgorithmStateChange_EventPatternParameters(options *SageMakerAlgorithmStateChange_SageMakerAlgorithmStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

