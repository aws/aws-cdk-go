//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func validateOperatingSystemFamily_OfParameters(family *string) error {
	return nil
}

