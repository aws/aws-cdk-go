//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"
)

func validateIdentitySource_ContextParameters(context *string) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func validateIdentitySource_HeaderParameters(headerName *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	return nil
}

func validateIdentitySource_QueryStringParameters(queryString *string) error {
	if queryString == nil {
		return fmt.Errorf("parameter queryString is required, but nil was provided")
	}

	return nil
}

func validateIdentitySource_StageVariableParameters(stageVariable *string) error {
	if stageVariable == nil {
		return fmt.Errorf("parameter stageVariable is required, but nil was provided")
	}

	return nil
}

