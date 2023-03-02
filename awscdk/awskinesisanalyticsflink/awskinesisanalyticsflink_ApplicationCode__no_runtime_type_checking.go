//go:build no_runtime_type_checking

package awskinesisanalyticsflink

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationCode) validateBindParameters(scope awscdk.Construct) error {
	return nil
}

func validateApplicationCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateApplicationCode_FromBucketParameters(bucket awss3.IBucket, fileKey *string) error {
	return nil
}

