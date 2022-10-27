//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaIntegration) validateBindParameters(method Method) error {
	return nil
}

func validateNewLambdaIntegrationParameters(handler awslambda.IFunction, options *LambdaIntegrationOptions) error {
	return nil
}

