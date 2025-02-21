//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func validateInlineImportSource_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateInlineImportSource_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateInlineImportSource_FromInlineParameters(data *string) error {
	return nil
}

func validateNewInlineImportSourceParameters(data *string) error {
	return nil
}

