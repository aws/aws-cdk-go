//go:build !no_runtime_type_checking

package awscodedeploy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TimeBasedCanaryTrafficRouting) validateBindParameters(_scope constructs.Construct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateTimeBasedCanaryTrafficRouting_TimeBasedCanaryParameters(props *TimeBasedCanaryTrafficRoutingProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateTimeBasedCanaryTrafficRouting_TimeBasedLinearParameters(props *TimeBasedLinearTrafficRoutingProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewTimeBasedCanaryTrafficRoutingParameters(props *TimeBasedCanaryTrafficRoutingProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

