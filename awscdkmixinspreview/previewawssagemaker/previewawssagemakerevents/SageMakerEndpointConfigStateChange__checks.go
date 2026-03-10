//go:build !no_runtime_type_checking

package previewawssagemakerevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateSageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangePatternParameters(options *SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

