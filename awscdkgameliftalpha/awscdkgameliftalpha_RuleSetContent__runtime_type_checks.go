//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (r *jsiiProxy_RuleSetContent) validateBindParameters(_scope constructs.Construct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateRuleSetContent_FromInlineParameters(body *string) error {
	if body == nil {
		return fmt.Errorf("parameter body is required, but nil was provided")
	}

	return nil
}

func validateRuleSetContent_FromJsonFileParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateNewRuleSetContentParameters(props *RuleSetContentProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

