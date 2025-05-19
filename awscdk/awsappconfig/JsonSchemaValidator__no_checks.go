//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func validateJsonSchemaValidator_FromFileParameters(inputPath *string) error {
	return nil
}

func validateJsonSchemaValidator_FromInlineParameters(code *string) error {
	return nil
}

