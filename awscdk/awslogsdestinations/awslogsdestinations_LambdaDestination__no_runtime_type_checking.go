//go:build no_runtime_type_checking

package awslogsdestinations

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaDestination) validateBindParameters(scope awscdk.Construct, logGroup awslogs.ILogGroup) error {
	return nil
}

func validateNewLambdaDestinationParameters(fn awslambda.IFunction, options *LambdaDestinationOptions) error {
	return nil
}

