//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebSocketRouteIntegration) validateBindParameters(options *WebSocketRouteIntegrationBindOptions) error {
	return nil
}

func validateNewWebSocketRouteIntegrationParameters(id *string) error {
	return nil
}

