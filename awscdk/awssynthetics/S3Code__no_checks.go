//go:build no_runtime_type_checking

package awssynthetics

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Code) validateBindParameters(_scope constructs.Construct, _handler *string, _family RuntimeFamily) error {
	return nil
}

func validateS3Code_FromAssetParameters(assetPath *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3Code_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateS3Code_FromInlineParameters(code *string) error {
	return nil
}

func validateNewS3CodeParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

