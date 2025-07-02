//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaFunction) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (l *jsiiProxy_LambdaFunction) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewLambdaFunctionParameters(lambdaFunction awslambda.IFunction, parameters *LambdaFunctionParameters) error {
	return nil
}

