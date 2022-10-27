//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (r *jsiiProxy_Runtime) validateRuntimeEqualsParameters(other Runtime) error {
	if other == nil {
		return fmt.Errorf("parameter other is required, but nil was provided")
	}

	return nil
}

func validateNewRuntimeParameters(name *string, props *LambdaRuntimeProps) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

