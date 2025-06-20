//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func validateDefaultValue_ValueParameters(value interface{}) error {
	return nil
}

