//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (m *jsiiProxy_MockIntegration) validateBindParameters(method Method) error {
	if method == nil {
		return fmt.Errorf("parameter method is required, but nil was provided")
	}

	return nil
}

func validateNewMockIntegrationParameters(options *IntegrationOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

