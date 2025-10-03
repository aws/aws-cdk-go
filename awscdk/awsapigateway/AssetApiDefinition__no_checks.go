//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetApiDefinition) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (a *jsiiProxy_AssetApiDefinition) validateBindAfterCreateParameters(scope constructs.Construct, restApi IRestApi) error {
	return nil
}

func validateAssetApiDefinition_FromAssetParameters(file *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateAssetApiDefinition_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateAssetApiDefinition_FromInlineParameters(definition interface{}) error {
	return nil
}

func validateNewAssetApiDefinitionParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

