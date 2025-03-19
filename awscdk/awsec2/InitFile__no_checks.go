//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateInitFile_FromAssetParameters(targetFileName *string, path *string, options *InitFileAssetOptions) error {
	return nil
}

func validateInitFile_FromExistingAssetParameters(targetFileName *string, asset awss3assets.Asset, options *InitFileOptions) error {
	return nil
}

func validateInitFile_FromFileInlineParameters(targetFileName *string, sourceFileName *string, options *InitFileOptions) error {
	return nil
}

func validateInitFile_FromObjectParameters(fileName *string, obj *map[string]interface{}, options *InitFileOptions) error {
	return nil
}

func validateInitFile_FromS3ObjectParameters(fileName *string, bucket awss3.IBucket, key *string, options *InitFileOptions) error {
	return nil
}

func validateInitFile_FromStringParameters(fileName *string, content *string, options *InitFileOptions) error {
	return nil
}

func validateInitFile_FromUrlParameters(fileName *string, url *string, options *InitFileOptions) error {
	return nil
}

func validateInitFile_SymlinkParameters(fileName *string, target *string, options *InitFileOptions) error {
	return nil
}

func validateNewInitFileParameters(fileName *string, options *InitFileOptions) error {
	return nil
}

