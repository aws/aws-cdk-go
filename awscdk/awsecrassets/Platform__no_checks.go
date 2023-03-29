//go:build no_runtime_type_checking

package awsecrassets

// Building without runtime type checking enabled, so all the below just return nil

func validatePlatform_CustomParameters(platform *string) error {
	return nil
}

