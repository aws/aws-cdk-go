//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateS3ApiSchema_FromInlineParameters(schema *string) error {
	return nil
}

func validateS3ApiSchema_FromLocalAssetParameters(path *string) error {
	return nil
}

func validateS3ApiSchema_FromS3FileParameters(bucket awss3.IBucketRef, objectKey *string) error {
	return nil
}

func validateNewS3ApiSchemaParameters(location *awss3.Location) error {
	return nil
}

