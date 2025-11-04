//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3OnFailureDestination) validateBindParameters(target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	return nil
}

func validateNewS3OnFailureDestinationParameters(bucket awss3.IBucket) error {
	return nil
}

