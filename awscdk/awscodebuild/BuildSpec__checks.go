//go:build !no_runtime_type_checking

package awscodebuild

import (
	"fmt"
)

func validateBuildSpec_FromObjectParameters(value *map[string]interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateBuildSpec_FromObjectToYamlParameters(value *map[string]interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateBuildSpec_FromSourceFilenameParameters(filename *string) error {
	if filename == nil {
		return fmt.Errorf("parameter filename is required, but nil was provided")
	}

	return nil
}

