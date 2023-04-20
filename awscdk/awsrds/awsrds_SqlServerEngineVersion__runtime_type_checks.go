//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func validateSqlServerEngineVersion_OfParameters(sqlServerFullVersion *string, sqlServerMajorVersion *string) error {
	if sqlServerFullVersion == nil {
		return fmt.Errorf("parameter sqlServerFullVersion is required, but nil was provided")
	}

	if sqlServerMajorVersion == nil {
		return fmt.Errorf("parameter sqlServerMajorVersion is required, but nil was provided")
	}

	return nil
}

