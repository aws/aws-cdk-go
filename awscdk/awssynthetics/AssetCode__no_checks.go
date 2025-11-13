//go:build no_runtime_type_checking

package awssynthetics

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetCode) validateBindParameters(scope constructs.Construct, handler *string, family RuntimeFamily) error {
	return nil
}

func validateAssetCode_FromAssetParameters(assetPath *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateAssetCode_FromBucketParameters(bucket interfacesawss3.IBucketRef, key *string) error {
	return nil
}

func validateAssetCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewAssetCodeParameters(assetPath *string, options *awss3assets.AssetOptions) error {
	return nil
}

