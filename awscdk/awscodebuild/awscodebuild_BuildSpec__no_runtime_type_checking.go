//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func validateBuildSpec_FromObjectParameters(value *map[string]interface{}) error {
	return nil
}

func validateBuildSpec_FromObjectToYamlParameters(value *map[string]interface{}) error {
	return nil
}

func validateBuildSpec_FromSourceFilenameParameters(filename *string) error {
	return nil
}

