//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateApiSchema_FromInlineParameters(schema *string) error {
	return nil
}

func validateApiSchema_FromLocalAssetParameters(path *string) error {
	return nil
}

func validateApiSchema_FromS3FileParameters(bucket interfacesawss3.IBucketRef, objectKey *string) error {
	return nil
}

func validateNewApiSchemaParameters(s3File *awss3.Location) error {
	return nil
}

