//go:build no_runtime_type_checking

package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetContent) validateBindParameters(scope constructs.Construct, role awsiam.IRole) error {
	return nil
}

func validateAssetContent_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateAssetContent_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewAssetContentParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

