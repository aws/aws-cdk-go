//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func validateS3ImportSource_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3ImportSource_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateS3ImportSource_FromInlineParameters(data *string) error {
	return nil
}

func validateNewS3ImportSourceParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

