//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetEnvironmentFile) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateAssetEnvironmentFile_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateAssetEnvironmentFile_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewAssetEnvironmentFileParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

