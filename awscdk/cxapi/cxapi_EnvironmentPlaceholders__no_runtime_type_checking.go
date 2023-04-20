//go:build no_runtime_type_checking

package cxapi

// Building without runtime type checking enabled, so all the below just return nil

func validateEnvironmentPlaceholders_ReplaceParameters(object interface{}, values *EnvironmentPlaceholderValues) error {
	return nil
}

func validateEnvironmentPlaceholders_ReplaceAsyncParameters(object interface{}, provider IEnvironmentPlaceholderProvider) error {
	return nil
}

