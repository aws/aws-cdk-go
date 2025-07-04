//go:build !no_runtime_type_checking

package awscdkiotactionsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

func validateNewStepFunctionsStateMachineActionParameters(stateMachine awsstepfunctions.IStateMachine, props *StepFunctionsStateMachineActionProps) error {
	if stateMachine == nil {
		return fmt.Errorf("parameter stateMachine is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

