//go:build no_runtime_type_checking

package awsapprunner

// Building without runtime type checking enabled, so all the below just return nil

func validateMemory_OfParameters(unit *string) error {
	return nil
}

