//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func validateSqlServerEngineVersion_OfParameters(sqlServerFullVersion *string, sqlServerMajorVersion *string) error {
	return nil
}

