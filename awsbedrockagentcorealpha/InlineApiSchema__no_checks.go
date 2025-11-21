//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineApiSchema) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_InlineApiSchema) validateGrantPermissionsToRoleParameters(role awsiam.IRole) error {
	return nil
}

func validateInlineApiSchema_FromInlineParameters(schema *string) error {
	return nil
}

func validateInlineApiSchema_FromLocalAssetParameters(path *string) error {
	return nil
}

func validateInlineApiSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	return nil
}

func validateNewInlineApiSchemaParameters(schema *string) error {
	return nil
}

