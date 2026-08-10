//go:build !no_runtime_type_checking

package previewawsstatesevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstates"
)

func (e *jsiiProxy_ExecutionEvents) validateStepFunctionsExecutionStatusChangePatternParameters(options *StepFunctionsExecutionStatusChange_StepFunctionsExecutionStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateExecutionEvents_FromExecutionParameters(executionRef interfacesawsstates.IExecutionRef) error {
	if executionRef == nil {
		return fmt.Errorf("parameter executionRef is required, but nil was provided")
	}

	return nil
}

