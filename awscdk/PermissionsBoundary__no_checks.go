//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validatePermissionsBoundary_FromArnParameters(arn *string) error {
	return nil
}

func validatePermissionsBoundary_FromNameParameters(name *string) error {
	return nil
}

