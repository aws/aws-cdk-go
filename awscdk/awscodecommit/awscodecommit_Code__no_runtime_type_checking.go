//go:build no_runtime_type_checking

package awscodecommit

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Code) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateCode_FromAssetParameters(asset awss3assets.Asset) error {
	return nil
}

func validateCode_FromDirectoryParameters(directoryPath *string) error {
	return nil
}

func validateCode_FromZipFileParameters(filePath *string) error {
	return nil
}

