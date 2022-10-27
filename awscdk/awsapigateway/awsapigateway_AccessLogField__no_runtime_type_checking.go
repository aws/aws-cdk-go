//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func validateAccessLogField_ContextAuthorizerParameters(property *string) error {
	return nil
}

func validateAccessLogField_ContextAuthorizerClaimsParameters(property *string) error {
	return nil
}

func validateAccessLogField_ContextRequestOverrideHeaderParameters(headerName *string) error {
	return nil
}

func validateAccessLogField_ContextRequestOverridePathParameters(pathName *string) error {
	return nil
}

func validateAccessLogField_ContextRequestOverrideQuerystringParameters(querystringName *string) error {
	return nil
}

func validateAccessLogField_ContextResponseOverrideHeaderParameters(headerName *string) error {
	return nil
}

