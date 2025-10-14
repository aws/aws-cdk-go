//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetApiSchema) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateAssetApiSchema_FromInlineParameters(schema *string) error {
	return nil
}

func validateAssetApiSchema_FromLocalAssetParameters(path *string) error {
	return nil
}

func validateAssetApiSchema_FromS3FileParameters(bucket awss3.IBucketRef, objectKey *string) error {
	return nil
}

func validateNewAssetApiSchemaParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

