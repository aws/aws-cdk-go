//go:build !no_runtime_type_checking

package awscdkintegtestsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (h *jsiiProxy_HttpApiCall) validateAssertAtPathParameters(_path *string, _expected ExpectedResult) error {
	if _path == nil {
		return fmt.Errorf("parameter _path is required, but nil was provided")
	}

	if _expected == nil {
		return fmt.Errorf("parameter _expected is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpApiCall) validateExpectParameters(expected ExpectedResult) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpApiCall) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpApiCall) validateGetAttStringParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpApiCall) validateNextParameters(next IApiCall) error {
	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HttpApiCall) validateWaitForAssertionsParameters(options *WaiterStateMachineOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHttpApiCall_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HttpApiCall) validateSetFlattenResponseParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewHttpApiCallParameters(scope constructs.Construct, id *string, props *HttpCallProps) error {
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

