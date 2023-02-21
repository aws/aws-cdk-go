//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_Integration) validateBindParameters(_method Method) error {
	if _method == nil {
		return fmt.Errorf("parameter _method is required, but nil was provided")
	}

	return nil
}

func validateNewIntegrationParameters(props *IntegrationProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

