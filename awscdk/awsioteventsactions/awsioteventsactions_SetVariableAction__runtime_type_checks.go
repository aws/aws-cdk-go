//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsioteventsactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiotevents"
	"github.com/aws/constructs-go/constructs/v3"
)

func (s *jsiiProxy_SetVariableAction) validateBindParameters(_scope constructs.Construct, _options *awsiotevents.ActionBindOptions) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func validateNewSetVariableActionParameters(variableName *string, value awsiotevents.Expression) error {
	if variableName == nil {
		return fmt.Errorf("parameter variableName is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

