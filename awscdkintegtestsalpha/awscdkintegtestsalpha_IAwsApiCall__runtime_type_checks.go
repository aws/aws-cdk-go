//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	"fmt"
)

func (i *jsiiProxy_IAwsApiCall) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IAwsApiCall) validateExpectParameters(expected ExpectedResult) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IAwsApiCall) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IAwsApiCall) validateGetAttStringParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

