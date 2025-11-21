//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_ToolSchema) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (t *jsiiProxy_ToolSchema) validateGrantPermissionsToRoleParameters(role awsiam.IRole) error {
	return nil
}

func validateToolSchema_FromInlineParameters(schema *[]*ToolDefinition) error {
	return nil
}

func validateToolSchema_FromLocalAssetParameters(path *string) error {
	return nil
}

func validateToolSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	return nil
}

func validateNewToolSchemaParameters(s3File *awss3.Location, inlineSchema *[]*ToolDefinition) error {
	return nil
}

