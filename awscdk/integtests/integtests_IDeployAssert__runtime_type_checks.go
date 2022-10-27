//go:build !no_runtime_type_checking

package integtests

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IDeployAssert) validateAwsApiCallParameters(service *string, api *string) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDeployAssert) validateExpectParameters(id *string, expected ExpectedResult, actual ActualResult) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	if actual == nil {
		return fmt.Errorf("parameter actual is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDeployAssert) validateInvokeFunctionParameters(props *LambdaInvokeFunctionProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

