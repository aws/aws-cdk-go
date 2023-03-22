//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func validateIdentitySource_ContextParameters(context *string) error {
	return nil
}

func validateIdentitySource_HeaderParameters(headerName *string) error {
	return nil
}

func validateIdentitySource_QueryStringParameters(queryString *string) error {
	return nil
}

func validateIdentitySource_StageVariableParameters(stageVariable *string) error {
	return nil
}

