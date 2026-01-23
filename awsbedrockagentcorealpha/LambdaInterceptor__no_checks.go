//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaInterceptor) validateBindParameters(scope constructs.Construct, gateway IGateway) error {
	return nil
}

func validateLambdaInterceptor_ForRequestParameters(lambdaFunction awslambda.IFunction, options *InterceptorOptions) error {
	return nil
}

func validateLambdaInterceptor_ForResponseParameters(lambdaFunction awslambda.IFunction, options *InterceptorOptions) error {
	return nil
}

