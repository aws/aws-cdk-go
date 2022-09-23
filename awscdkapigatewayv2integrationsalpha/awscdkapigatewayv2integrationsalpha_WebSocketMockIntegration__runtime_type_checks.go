//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
)

func (w *jsiiProxy_WebSocketMockIntegration) validateBindParameters(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewWebSocketMockIntegrationParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

