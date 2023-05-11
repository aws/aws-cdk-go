//go:build !no_runtime_type_checking

// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
)

func (h *jsiiProxy_HttpLambdaIntegration) validateBindParameters(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (h *jsiiProxy_HttpLambdaIntegration) validateCompleteBindParameters(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewHttpLambdaIntegrationParameters(id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if handler == nil {
		return fmt.Errorf("parameter handler is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

