//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Code) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateCode_FromInlineParameters(code *string) error {
	return nil
}

