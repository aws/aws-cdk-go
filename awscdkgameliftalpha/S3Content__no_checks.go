//go:build no_runtime_type_checking

package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Content) validateBindParameters(scope constructs.Construct, role awsiam.IRole) error {
	return nil
}

func validateS3Content_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3Content_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewS3ContentParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

