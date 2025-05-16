//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func validateConfigurationContent_FromFileParameters(inputPath *string) error {
	return nil
}

func validateConfigurationContent_FromInlineParameters(content *string) error {
	return nil
}

func validateConfigurationContent_FromInlineJsonParameters(content *string) error {
	return nil
}

func validateConfigurationContent_FromInlineTextParameters(content *string) error {
	return nil
}

func validateConfigurationContent_FromInlineYamlParameters(content *string) error {
	return nil
}

