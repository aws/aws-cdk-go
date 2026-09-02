//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

func validateDataQualityTargetTable_FromTableParameters(database interfacesawsglue.IDatabaseRef, table ITable) error {
	if database == nil {
		return fmt.Errorf("parameter database is required, but nil was provided")
	}

	if table == nil {
		return fmt.Errorf("parameter table is required, but nil was provided")
	}

	return nil
}

func validateDataQualityTargetTable_FromTableNameParameters(database interfacesawsglue.IDatabaseRef, tableName *string) error {
	if database == nil {
		return fmt.Errorf("parameter database is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	return nil
}

