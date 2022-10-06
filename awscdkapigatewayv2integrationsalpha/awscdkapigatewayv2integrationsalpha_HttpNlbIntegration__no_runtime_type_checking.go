//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpNlbIntegration) validateBindParameters(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpNlbIntegration) validateCompleteBindParameters(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetConnectionTypeParameters(val awscdkapigatewayv2alpha.HttpConnectionType) error {
	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetHttpMethodParameters(val awscdkapigatewayv2alpha.HttpMethod) error {
	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetIntegrationTypeParameters(val awscdkapigatewayv2alpha.HttpIntegrationType) error {
	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetPayloadFormatVersionParameters(val awscdkapigatewayv2alpha.PayloadFormatVersion) error {
	return nil
}

func validateNewHttpNlbIntegrationParameters(id *string, listener awselasticloadbalancingv2.INetworkListener, props *HttpNlbIntegrationProps) error {
	return nil
}

