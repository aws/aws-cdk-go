//go:build no_runtime_type_checking

package awss3deployment

// Building without runtime type checking enabled, so all the below just return nil

func validateSource_AssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateSource_BucketParameters(bucket awss3.IBucket, zipObjectKey *string) error {
	return nil
}

func validateSource_DataParameters(objectKey *string, data *string) error {
	return nil
}

func validateSource_JsonDataParameters(objectKey *string, obj interface{}) error {
	return nil
}

