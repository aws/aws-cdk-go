//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsapigatewayv2integrations

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
)

func (h *jsiiProxy_HttpUrlIntegration) validateBindParameters(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_arg, func() string { return "parameter _arg" }); err != nil {
		return err
	}

	return nil
}

func (h *jsiiProxy_HttpUrlIntegration) validateCompleteBindParameters(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func validateNewHttpUrlIntegrationParameters(id *string, url *string, props *HttpUrlIntegrationProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

