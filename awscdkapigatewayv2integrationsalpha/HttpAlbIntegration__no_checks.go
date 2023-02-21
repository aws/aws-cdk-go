//go:build no_runtime_type_checking

// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpAlbIntegration) validateBindParameters(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpAlbIntegration) validateCompleteBindParameters(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (j *jsiiProxy_HttpAlbIntegration) validateSetConnectionTypeParameters(val awscdkapigatewayv2alpha.HttpConnectionType) error {
	return nil
}

func (j *jsiiProxy_HttpAlbIntegration) validateSetHttpMethodParameters(val awscdkapigatewayv2alpha.HttpMethod) error {
	return nil
}

func (j *jsiiProxy_HttpAlbIntegration) validateSetIntegrationTypeParameters(val awscdkapigatewayv2alpha.HttpIntegrationType) error {
	return nil
}

func (j *jsiiProxy_HttpAlbIntegration) validateSetPayloadFormatVersionParameters(val awscdkapigatewayv2alpha.PayloadFormatVersion) error {
	return nil
}

func validateNewHttpAlbIntegrationParameters(id *string, listener awselasticloadbalancingv2.IApplicationListener, props *HttpAlbIntegrationProps) error {
	return nil
}

