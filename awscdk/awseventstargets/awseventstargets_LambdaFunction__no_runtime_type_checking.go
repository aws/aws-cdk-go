//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaFunction) validateBindParameters(rule awsevents.IRule) error {
	return nil
}

func validateNewLambdaFunctionParameters(handler awslambda.IFunction, props *LambdaFunctionProps) error {
	return nil
}

