//go:build !no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AssertionsProvider) validateAddPolicyStatementFromSdkCallParameters(service *string, api *string) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateAddToRolePolicyParameters(statement interface{}) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateEncodeParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateGrantInvokeParameters(principalArn *string) error {
	if principalArn == nil {
		return fmt.Errorf("parameter principalArn is required, but nil was provided")
	}

	return nil
}

func validateAssertionsProvider_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewAssertionsProviderParameters(scope constructs.Construct, id *string, props *AssertionsProviderProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

