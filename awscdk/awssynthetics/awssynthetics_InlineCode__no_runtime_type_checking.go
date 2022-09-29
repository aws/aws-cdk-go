//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awssynthetics

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineCode) validateBindParameters(_scope constructs.Construct, handler *string, _family RuntimeFamily) error {
	return nil
}

func validateInlineCode_FromAssetParameters(assetPath *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateInlineCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateInlineCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewInlineCodeParameters(code *string) error {
	return nil
}

