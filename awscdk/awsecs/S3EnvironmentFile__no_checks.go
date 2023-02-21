//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3EnvironmentFile) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func validateS3EnvironmentFile_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3EnvironmentFile_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewS3EnvironmentFileParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

