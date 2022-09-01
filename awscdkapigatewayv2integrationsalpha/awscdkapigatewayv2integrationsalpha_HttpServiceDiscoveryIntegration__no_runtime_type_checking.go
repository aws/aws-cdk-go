//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) validateBindParameters(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) validateCompleteBindParameters(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetConnectionTypeParameters(val awscdkapigatewayv2alpha.HttpConnectionType) error {
	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetHttpMethodParameters(val awscdkapigatewayv2alpha.HttpMethod) error {
	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetIntegrationTypeParameters(val awscdkapigatewayv2alpha.HttpIntegrationType) error {
	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetPayloadFormatVersionParameters(val awscdkapigatewayv2alpha.PayloadFormatVersion) error {
	return nil
}

func validateNewHttpServiceDiscoveryIntegrationParameters(id *string, service awsservicediscovery.IService, props *HttpServiceDiscoveryIntegrationProps) error {
	return nil
}

