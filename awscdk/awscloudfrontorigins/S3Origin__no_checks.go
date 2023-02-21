//go:build no_runtime_type_checking

package awscloudfrontorigins

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Origin) validateBindParameters(scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	return nil
}

func validateNewS3OriginParameters(bucket awss3.IBucket, props *S3OriginProps) error {
	return nil
}

