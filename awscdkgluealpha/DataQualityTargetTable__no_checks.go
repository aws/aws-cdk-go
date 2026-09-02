//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateDataQualityTargetTable_FromTableParameters(database interfacesawsglue.IDatabaseRef, table ITable) error {
	return nil
}

func validateDataQualityTargetTable_FromTableNameParameters(database interfacesawsglue.IDatabaseRef, tableName *string) error {
	return nil
}

