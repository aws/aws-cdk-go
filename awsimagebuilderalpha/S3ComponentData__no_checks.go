//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3ComponentData) validateGrantPutParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3ComponentData) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateS3ComponentData_FromAssetParameters(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3ComponentData_FromComponentDocumentJsonObjectParameters(data *ComponentDocument) error {
	return nil
}

func validateS3ComponentData_FromInlineParameters(data *string) error {
	return nil
}

func validateS3ComponentData_FromJsonObjectParameters(data *map[string]interface{}) error {
	return nil
}

func validateS3ComponentData_FromS3Parameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewS3ComponentDataParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

