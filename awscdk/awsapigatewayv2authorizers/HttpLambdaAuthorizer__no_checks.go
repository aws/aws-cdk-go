//go:build no_runtime_type_checking

package awsapigatewayv2authorizers

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpLambdaAuthorizer) validateBindParameters(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) error {
	return nil
}

func validateNewHttpLambdaAuthorizerParameters(id *string, handler awslambda.IFunction, props *HttpLambdaAuthorizerProps) error {
	return nil
}

