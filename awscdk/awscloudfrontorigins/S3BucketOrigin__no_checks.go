//go:build no_runtime_type_checking

package awscloudfrontorigins

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketOrigin) validateBindParameters(scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	return nil
}

func validateS3BucketOrigin_WithBucketDefaultsParameters(bucket awss3.IBucket, props *awscloudfront.OriginProps) error {
	return nil
}

func validateS3BucketOrigin_WithOriginAccessControlParameters(bucket awss3.IBucket, props *S3BucketOriginWithOACProps) error {
	return nil
}

func validateS3BucketOrigin_WithOriginAccessIdentityParameters(bucket awss3.IBucket, props *S3BucketOriginWithOAIProps) error {
	return nil
}

func validateNewS3BucketOriginParameters(bucket awss3.IBucket, props *S3BucketOriginBaseProps) error {
	return nil
}

