//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateRuntimeCustomClaim_WithStringArrayValueParameters(name *string, values *[]*string) error {
	return nil
}

func validateRuntimeCustomClaim_WithStringValueParameters(name *string, value *string) error {
	return nil
}

