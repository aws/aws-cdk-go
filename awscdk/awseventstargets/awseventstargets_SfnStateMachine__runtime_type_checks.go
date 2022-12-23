//go:build !no_runtime_type_checking

package awseventstargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

func (s *jsiiProxy_SfnStateMachine) validateBindParameters(_rule awsevents.IRule) error {
	if _rule == nil {
		return fmt.Errorf("parameter _rule is required, but nil was provided")
	}

	return nil
}

func validateNewSfnStateMachineParameters(machine awsstepfunctions.IStateMachine, props *SfnStateMachineProps) error {
	if machine == nil {
		return fmt.Errorf("parameter machine is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

