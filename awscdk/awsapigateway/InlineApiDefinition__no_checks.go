//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineApiDefinition) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_InlineApiDefinition) validateBindAfterCreateParameters(_scope constructs.Construct, _restApi IRestApi) error {
	return nil
}

func validateInlineApiDefinition_FromAssetParameters(file *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateInlineApiDefinition_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateInlineApiDefinition_FromInlineParameters(definition interface{}) error {
	return nil
}

func validateNewInlineApiDefinitionParameters(definition interface{}) error {
	return nil
}

