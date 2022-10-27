//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3EventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func validateNewS3EventSourceParameters(bucket awss3.Bucket, props *S3EventSourceProps) error {
	return nil
}

