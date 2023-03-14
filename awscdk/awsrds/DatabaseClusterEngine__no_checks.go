//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func validateDatabaseClusterEngine_AuroraParameters(props *AuroraClusterEngineProps) error {
	return nil
}

func validateDatabaseClusterEngine_AuroraMysqlParameters(props *AuroraMysqlClusterEngineProps) error {
	return nil
}

func validateDatabaseClusterEngine_AuroraPostgresParameters(props *AuroraPostgresClusterEngineProps) error {
	return nil
}

