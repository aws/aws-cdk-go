//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAgentActionGroup_CodeInterpreterParameters(enabled *bool) error {
	if enabled == nil {
		return fmt.Errorf("parameter enabled is required, but nil was provided")
	}

	return nil
}

func validateAgentActionGroup_UserInputParameters(enabled *bool) error {
	if enabled == nil {
		return fmt.Errorf("parameter enabled is required, but nil was provided")
	}

	return nil
}

func validateNewAgentActionGroupParameters(props *AgentActionGroupProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

