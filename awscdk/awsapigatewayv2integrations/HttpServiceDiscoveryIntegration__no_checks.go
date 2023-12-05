//go:build no_runtime_type_checking

package awsapigatewayv2integrations

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) validateBindParameters(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) validateCompleteBindParameters(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetConnectionTypeParameters(val awsapigatewayv2.HttpConnectionType) error {
	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetHttpMethodParameters(val awsapigatewayv2.HttpMethod) error {
	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetIntegrationTypeParameters(val awsapigatewayv2.HttpIntegrationType) error {
	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetPayloadFormatVersionParameters(val awsapigatewayv2.PayloadFormatVersion) error {
	return nil
}

func validateNewHttpServiceDiscoveryIntegrationParameters(id *string, service awsservicediscovery.IService, props *HttpServiceDiscoveryIntegrationProps) error {
	return nil
}

