//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockFoundationModel) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BedrockFoundationModel) validateGrantInvokeAllRegionsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateBedrockFoundationModel_FromCdkFoundationModelParameters(modelId awsbedrock.FoundationModel, props *BedrockFoundationModelProps) error {
	return nil
}

func validateBedrockFoundationModel_FromCdkFoundationModelIdParameters(modelId awsbedrock.FoundationModelIdentifier, props *BedrockFoundationModelProps) error {
	return nil
}

func validateNewBedrockFoundationModelParameters(value *string, props *BedrockFoundationModelProps) error {
	return nil
}

