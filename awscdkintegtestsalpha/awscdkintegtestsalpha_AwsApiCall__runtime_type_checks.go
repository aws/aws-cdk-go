//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AwsApiCall) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsApiCall) validateExpectParameters(expected ExpectedResult) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsApiCall) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsApiCall) validateGetAttStringParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsApiCall) validateNextParameters(next IApiCall) error {
	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsApiCall) validateWaitForAssertionsParameters(options *WaiterStateMachineOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAwsApiCall_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsApiCall) validateSetFlattenResponseParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsApiCallParameters(scope constructs.Construct, id *string, props *AwsApiCallProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

