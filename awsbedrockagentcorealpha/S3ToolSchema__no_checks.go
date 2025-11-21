//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3ToolSchema) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (s *jsiiProxy_S3ToolSchema) validateGrantPermissionsToRoleParameters(role awsiam.IRole) error {
	return nil
}

func validateS3ToolSchema_FromInlineParameters(schema *[]*ToolDefinition) error {
	return nil
}

func validateS3ToolSchema_FromLocalAssetParameters(path *string) error {
	return nil
}

func validateS3ToolSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	return nil
}

func validateNewS3ToolSchemaParameters(location *awss3.Location) error {
	return nil
}

