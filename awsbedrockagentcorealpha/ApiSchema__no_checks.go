//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiSchema) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (a *jsiiProxy_ApiSchema) validateGrantPermissionsToRoleParameters(role awsiam.IRole) error {
	return nil
}

func validateApiSchema_FromInlineParameters(schema *string) error {
	return nil
}

func validateApiSchema_FromLocalAssetParameters(path *string) error {
	return nil
}

func validateApiSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	return nil
}

func validateNewApiSchemaParameters(s3File *awss3.Location) error {
	return nil
}

