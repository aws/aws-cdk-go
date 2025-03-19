//go:build no_runtime_type_checking

package awsapigatewayv2integrations

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebSocketAwsIntegration) validateBindParameters(_options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) error {
	return nil
}

func validateNewWebSocketAwsIntegrationParameters(id *string, props *WebSocketAwsIntegrationProps) error {
	return nil
}

