//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (a *jsiiProxy_AwsIntegration) validateBindParameters(method Method) error {
	if method == nil {
		return fmt.Errorf("parameter method is required, but nil was provided")
	}

	return nil
}

func validateNewAwsIntegrationParameters(props *AwsIntegrationProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

