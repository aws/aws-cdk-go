//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateKafkaSchemaValidationAttribute_OfParameters(name *string) error {
	return nil
}

