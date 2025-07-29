//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EnvironmentFile) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateEnvironmentFile_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateEnvironmentFile_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

