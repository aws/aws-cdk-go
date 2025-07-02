//go:build !no_runtime_type_checking

package cxapi

import (
	"fmt"
)

func validateEnvironmentUtils_FormatParameters(account *string, region *string) error {
	if account == nil {
		return fmt.Errorf("parameter account is required, but nil was provided")
	}

	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	return nil
}

func validateEnvironmentUtils_MakeParameters(account *string, region *string) error {
	if account == nil {
		return fmt.Errorf("parameter account is required, but nil was provided")
	}

	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	return nil
}

func validateEnvironmentUtils_ParseParameters(environment *string) error {
	if environment == nil {
		return fmt.Errorf("parameter environment is required, but nil was provided")
	}

	return nil
}

