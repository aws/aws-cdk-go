//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

func (s *jsiiProxy_StartExecution) validateBindParameters(task awsstepfunctions.Task) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

func validateNewStartExecutionParameters(stateMachine awsstepfunctions.IStateMachine, props *StartExecutionProps) error {
	if stateMachine == nil {
		return fmt.Errorf("parameter stateMachine is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

