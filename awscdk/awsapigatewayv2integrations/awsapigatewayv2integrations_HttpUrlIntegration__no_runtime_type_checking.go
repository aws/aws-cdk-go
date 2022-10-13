//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigatewayv2integrations

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpUrlIntegration) validateBindParameters(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpUrlIntegration) validateCompleteBindParameters(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func validateNewHttpUrlIntegrationParameters(id *string, url *string, props *HttpUrlIntegrationProps) error {
	return nil
}

