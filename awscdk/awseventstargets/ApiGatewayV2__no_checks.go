//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiGatewayV2) validateBindParameters(rule awsevents.IRule) error {
	return nil
}

func validateNewApiGatewayV2Parameters(httpApi awsapigatewayv2.IHttpApi, props *ApiGatewayProps) error {
	return nil
}

