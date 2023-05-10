//go:build no_runtime_type_checking

// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpLambdaIntegration) validateBindParameters(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpLambdaIntegration) validateCompleteBindParameters(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func validateNewHttpLambdaIntegrationParameters(id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) error {
	return nil
}

