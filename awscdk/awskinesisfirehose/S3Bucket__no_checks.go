//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Bucket) validateBindParameters(scope constructs.Construct, options *DestinationBindOptions) error {
	return nil
}

func validateNewS3BucketParameters(bucket awss3.IBucket, props *S3BucketProps) error {
	return nil
}

