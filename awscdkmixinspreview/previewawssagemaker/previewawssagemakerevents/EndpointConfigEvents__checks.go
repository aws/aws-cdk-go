//go:build !no_runtime_type_checking

package previewawssagemakerevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker"
)

func (e *jsiiProxy_EndpointConfigEvents) validateSageMakerEndpointConfigStateChangePatternParameters(options *EndpointConfigEvents_SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEndpointConfigEvents_FromEndpointConfigParameters(endpointConfigRef interfacesawssagemaker.IEndpointConfigRef) error {
	if endpointConfigRef == nil {
		return fmt.Errorf("parameter endpointConfigRef is required, but nil was provided")
	}

	return nil
}

