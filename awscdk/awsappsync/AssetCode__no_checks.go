//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateAssetCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateAssetCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewAssetCodeParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

