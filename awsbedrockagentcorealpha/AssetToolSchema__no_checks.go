//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetToolSchema) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (a *jsiiProxy_AssetToolSchema) validateGrantPermissionsToRoleParameters(role awsiam.IRole) error {
	return nil
}

func validateAssetToolSchema_FromInlineParameters(schema *[]*ToolDefinition) error {
	return nil
}

func validateAssetToolSchema_FromLocalAssetParameters(path *string) error {
	return nil
}

func validateAssetToolSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	return nil
}

func validateNewAssetToolSchemaParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

