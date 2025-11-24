//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateDockerfileData_FromAssetParameters(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateDockerfileData_FromInlineParameters(data *string) error {
	return nil
}

func validateDockerfileData_FromS3Parameters(bucket awss3.IBucket, key *string) error {
	return nil
}

