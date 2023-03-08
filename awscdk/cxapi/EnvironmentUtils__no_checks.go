//go:build no_runtime_type_checking

package cxapi

// Building without runtime type checking enabled, so all the below just return nil

func validateEnvironmentUtils_FormatParameters(account *string, region *string) error {
	return nil
}

func validateEnvironmentUtils_MakeParameters(account *string, region *string) error {
	return nil
}

func validateEnvironmentUtils_ParseParameters(environment *string) error {
	return nil
}

