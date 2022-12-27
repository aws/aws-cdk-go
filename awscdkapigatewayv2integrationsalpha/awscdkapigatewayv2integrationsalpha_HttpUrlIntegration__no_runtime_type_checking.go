//go:build no_runtime_type_checking

// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpUrlIntegration) validateBindParameters(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpUrlIntegration) validateCompleteBindParameters(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func validateNewHttpUrlIntegrationParameters(id *string, url *string, props *HttpUrlIntegrationProps) error {
	return nil
}

