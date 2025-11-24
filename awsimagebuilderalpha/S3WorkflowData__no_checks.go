//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3WorkflowData) validateGrantPutParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3WorkflowData) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateS3WorkflowData_FromAssetParameters(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3WorkflowData_FromInlineParameters(data *string) error {
	return nil
}

func validateS3WorkflowData_FromJsonObjectParameters(data *map[string]interface{}) error {
	return nil
}

func validateS3WorkflowData_FromS3Parameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewS3WorkflowDataParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

