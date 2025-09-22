//go:build no_runtime_type_checking

package awscodepipeline

// Building without runtime type checking enabled, so all the below just return nil

func validateCustomActionRegistration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCustomActionRegistrationParameters(scope constructs.Construct, id *string, props *CustomActionRegistrationProps) error {
	return nil
}

