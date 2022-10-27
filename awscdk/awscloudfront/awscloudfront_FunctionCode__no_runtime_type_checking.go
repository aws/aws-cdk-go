//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func validateFunctionCode_FromFileParameters(options *FileCodeOptions) error {
	return nil
}

func validateFunctionCode_FromInlineParameters(code *string) error {
	return nil
}

