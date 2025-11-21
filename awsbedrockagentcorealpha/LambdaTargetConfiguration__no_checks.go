//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaTargetConfiguration) validateBindParameters(scope constructs.Construct, gateway IGateway) error {
	return nil
}

func validateLambdaTargetConfiguration_CreateParameters(lambdaFunction awslambda.IFunction, toolSchema ToolSchema) error {
	return nil
}

func validateNewLambdaTargetConfigurationParameters(lambdaFunction awslambda.IFunction, toolSchema ToolSchema) error {
	return nil
}

