//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3EventSourceV2) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func validateNewS3EventSourceV2Parameters(bucket awss3.IBucket, props *S3EventSourceProps) error {
	return nil
}

