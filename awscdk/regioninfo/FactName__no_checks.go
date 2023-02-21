//go:build no_runtime_type_checking

package regioninfo

// Building without runtime type checking enabled, so all the below just return nil

func validateFactName_AdotLambdaLayerParameters(type_ *string, version *string, architecture *string) error {
	return nil
}

func validateFactName_CloudwatchLambdaInsightsVersionParameters(version *string) error {
	return nil
}

func validateFactName_ServicePrincipalParameters(service *string) error {
	return nil
}

