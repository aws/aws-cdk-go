//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Code) validateBindParameters(scope constructs.Construct, grantable awsiam.IGrantable) error {
	return nil
}

func validateCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

