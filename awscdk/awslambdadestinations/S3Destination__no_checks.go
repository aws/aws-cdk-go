//go:build no_runtime_type_checking

package awslambdadestinations

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Destination) validateBindParameters(scope constructs.Construct, fn awslambda.IFunction, options *awslambda.DestinationOptions) error {
	return nil
}

func validateNewS3DestinationParameters(bucket awss3.IBucket) error {
	return nil
}

