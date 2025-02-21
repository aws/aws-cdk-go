//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func validateMysqlEngineVersion_OfParameters(mysqlFullVersion *string, mysqlMajorVersion *string) error {
	if mysqlFullVersion == nil {
		return fmt.Errorf("parameter mysqlFullVersion is required, but nil was provided")
	}

	if mysqlMajorVersion == nil {
		return fmt.Errorf("parameter mysqlMajorVersion is required, but nil was provided")
	}

	return nil
}

