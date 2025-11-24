//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3DockerfileData) validateGrantPutParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3DockerfileData) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateS3DockerfileData_FromAssetParameters(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3DockerfileData_FromInlineParameters(data *string) error {
	return nil
}

func validateS3DockerfileData_FromS3Parameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewS3DockerfileDataParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

