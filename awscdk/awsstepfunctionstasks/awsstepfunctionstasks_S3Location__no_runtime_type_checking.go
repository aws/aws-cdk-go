//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Location) validateBindParameters(task ISageMakerTask, opts *S3LocationBindOptions) error {
	return nil
}

func validateS3Location_FromBucketParameters(bucket awss3.IBucket, keyPrefix *string) error {
	return nil
}

func validateS3Location_FromJsonExpressionParameters(expression *string) error {
	return nil
}

