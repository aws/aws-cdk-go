//go:build no_runtime_type_checking

package awss3notifications

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaDestination) validateBindParameters(scope constructs.Construct, bucket awss3.IBucketRef) error {
	return nil
}

func validateNewLambdaDestinationParameters(fn awslambda.IFunction) error {
	return nil
}

