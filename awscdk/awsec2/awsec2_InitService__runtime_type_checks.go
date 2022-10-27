//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateInitService_DisableParameters(serviceName *string) error {
	if serviceName == nil {
		return fmt.Errorf("parameter serviceName is required, but nil was provided")
	}

	return nil
}

func validateInitService_EnableParameters(serviceName *string, options *InitServiceOptions) error {
	if serviceName == nil {
		return fmt.Errorf("parameter serviceName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

