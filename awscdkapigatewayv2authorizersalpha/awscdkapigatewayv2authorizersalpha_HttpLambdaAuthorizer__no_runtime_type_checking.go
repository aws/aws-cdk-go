//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Authorizers for AWS APIGateway V2
package awscdkapigatewayv2authorizersalpha

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpLambdaAuthorizer) validateBindParameters(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) error {
	return nil
}

func validateNewHttpLambdaAuthorizerParameters(id *string, handler awslambda.IFunction, props *HttpLambdaAuthorizerProps) error {
	return nil
}

