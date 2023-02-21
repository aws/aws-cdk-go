//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiDefinition) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (a *jsiiProxy_ApiDefinition) validateBindAfterCreateParameters(_scope constructs.Construct, _restApi IRestApi) error {
	return nil
}

func validateApiDefinition_FromAssetParameters(file *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateApiDefinition_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateApiDefinition_FromInlineParameters(definition interface{}) error {
	return nil
}

