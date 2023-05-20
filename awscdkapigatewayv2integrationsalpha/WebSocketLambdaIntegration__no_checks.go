//go:build no_runtime_type_checking

package awscdkapigatewayv2integrationsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebSocketLambdaIntegration) validateBindParameters(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) error {
	return nil
}

func validateNewWebSocketLambdaIntegrationParameters(id *string, handler awslambda.IFunction) error {
	return nil
}

