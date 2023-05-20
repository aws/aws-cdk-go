//go:build no_runtime_type_checking

package awscdkapigatewayv2authorizersalpha

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebSocketLambdaAuthorizer) validateBindParameters(options *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerBindOptions) error {
	return nil
}

func validateNewWebSocketLambdaAuthorizerParameters(id *string, handler awslambda.IFunction, props *WebSocketLambdaAuthorizerProps) error {
	return nil
}

