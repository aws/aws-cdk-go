//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"
)

func validateNewDataQualityTargetTableParameters(databaseName *string, tableName *string) error {
	if databaseName == nil {
		return fmt.Errorf("parameter databaseName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	return nil
}

