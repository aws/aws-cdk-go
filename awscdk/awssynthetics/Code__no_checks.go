//go:build no_runtime_type_checking

package awssynthetics

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Code) validateBindParameters(scope constructs.Construct, handler *string, family RuntimeFamily) error {
	return nil
}

func validateCode_FromAssetParameters(assetPath *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateCode_FromBucketParameters(bucket interfacesawss3.IBucketRef, key *string) error {
	return nil
}

func validateCode_FromInlineParameters(code *string) error {
	return nil
}

