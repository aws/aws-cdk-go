//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateInlineCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateInlineCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewInlineCodeParameters(code *string) error {
	return nil
}

