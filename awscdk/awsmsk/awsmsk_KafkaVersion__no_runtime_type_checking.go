//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsmsk

// Building without runtime type checking enabled, so all the below just return nil

func validateKafkaVersion_OfParameters(version *string) error {
	return nil
}

