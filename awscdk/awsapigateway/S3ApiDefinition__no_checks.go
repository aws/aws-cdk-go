//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3ApiDefinition) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func (s *jsiiProxy_S3ApiDefinition) validateBindAfterCreateParameters(_scope constructs.Construct, _restApi IRestApi) error {
	return nil
}

func validateS3ApiDefinition_FromAssetParameters(file *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3ApiDefinition_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateS3ApiDefinition_FromInlineParameters(definition interface{}) error {
	return nil
}

func validateNewS3ApiDefinitionParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

