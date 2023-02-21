//go:build !no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IApiCall) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApiCall) validateExpectParameters(expected ExpectedResult) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApiCall) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApiCall) validateGetAttStringParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApiCall) validateNextParameters(next IApiCall) error {
	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApiCall) validateWaitForAssertionsParameters(options *WaiterStateMachineOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

