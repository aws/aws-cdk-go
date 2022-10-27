//go:build no_runtime_type_checking

package awsapigatewayv2integrations

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpNlbIntegration) validateBindParameters(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpNlbIntegration) validateCompleteBindParameters(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetConnectionTypeParameters(val awsapigatewayv2.HttpConnectionType) error {
	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetHttpMethodParameters(val awsapigatewayv2.HttpMethod) error {
	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetIntegrationTypeParameters(val awsapigatewayv2.HttpIntegrationType) error {
	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetPayloadFormatVersionParameters(val awsapigatewayv2.PayloadFormatVersion) error {
	return nil
}

func validateNewHttpNlbIntegrationParameters(id *string, listener awselasticloadbalancingv2.INetworkListener, props *HttpNlbIntegrationProps) error {
	return nil
}

