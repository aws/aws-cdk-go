//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineToolSchema) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_InlineToolSchema) validateGrantPermissionsToRoleParameters(role awsiam.IRole) error {
	return nil
}

func validateInlineToolSchema_FromInlineParameters(schema *[]*ToolDefinition) error {
	return nil
}

func validateInlineToolSchema_FromLocalAssetParameters(path *string) error {
	return nil
}

func validateInlineToolSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	return nil
}

func validateNewInlineToolSchemaParameters(schema *[]*ToolDefinition) error {
	return nil
}

