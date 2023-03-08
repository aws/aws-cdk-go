//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func validatePostgresEngineVersion_OfParameters(postgresFullVersion *string, postgresMajorVersion *string, postgresFeatures *PostgresEngineFeatures) error {
	return nil
}

