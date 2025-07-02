//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func validateOracleEngineVersion_OfParameters(oracleFullVersion *string, oracleMajorVersion *string) error {
	if oracleFullVersion == nil {
		return fmt.Errorf("parameter oracleFullVersion is required, but nil was provided")
	}

	if oracleMajorVersion == nil {
		return fmt.Errorf("parameter oracleMajorVersion is required, but nil was provided")
	}

	return nil
}

