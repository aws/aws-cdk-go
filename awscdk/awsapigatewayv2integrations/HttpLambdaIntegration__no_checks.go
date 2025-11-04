//go:build no_runtime_type_checking

package awsapigatewayv2integrations

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpLambdaIntegration) validateBindParameters(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func (h *jsiiProxy_HttpLambdaIntegration) validateCompleteBindParameters(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	return nil
}

func validateNewHttpLambdaIntegrationParameters(id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) error {
	return nil
}

