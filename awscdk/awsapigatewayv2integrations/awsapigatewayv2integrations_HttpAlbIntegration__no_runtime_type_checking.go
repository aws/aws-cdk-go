//go:build no_runtime_type_checking

package awsapigatewayv2integrations

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpAlbIntegration) validateBindParameters(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpAlbIntegration) validateCompleteBindParameters(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (j *jsiiProxy_HttpAlbIntegration) validateSetConnectionTypeParameters(val awsapigatewayv2.HttpConnectionType) error {
	return nil
}

func (j *jsiiProxy_HttpAlbIntegration) validateSetHttpMethodParameters(val awsapigatewayv2.HttpMethod) error {
	return nil
}

func (j *jsiiProxy_HttpAlbIntegration) validateSetIntegrationTypeParameters(val awsapigatewayv2.HttpIntegrationType) error {
	return nil
}

func (j *jsiiProxy_HttpAlbIntegration) validateSetPayloadFormatVersionParameters(val awsapigatewayv2.PayloadFormatVersion) error {
	return nil
}

func validateNewHttpAlbIntegrationParameters(id *string, listener awselasticloadbalancingv2.IApplicationListener, props *HttpAlbIntegrationProps) error {
	return nil
}

