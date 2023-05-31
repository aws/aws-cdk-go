//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"
)

func validateAccessLogField_ContextAuthorizerParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func validateAccessLogField_ContextAuthorizerClaimsParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func validateAccessLogField_ContextRequestOverrideHeaderParameters(headerName *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	return nil
}

func validateAccessLogField_ContextRequestOverridePathParameters(pathName *string) error {
	if pathName == nil {
		return fmt.Errorf("parameter pathName is required, but nil was provided")
	}

	return nil
}

func validateAccessLogField_ContextRequestOverrideQuerystringParameters(querystringName *string) error {
	if querystringName == nil {
		return fmt.Errorf("parameter querystringName is required, but nil was provided")
	}

	return nil
}

func validateAccessLogField_ContextResponseOverrideHeaderParameters(headerName *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	return nil
}

