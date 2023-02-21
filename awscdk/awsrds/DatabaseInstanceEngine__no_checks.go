//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func validateDatabaseInstanceEngine_MariaDbParameters(props *MariaDbInstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_MysqlParameters(props *MySqlInstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_OracleEeParameters(props *OracleEeInstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_OracleEeCdbParameters(props *OracleEeCdbInstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_OracleSe2Parameters(props *OracleSe2InstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_OracleSe2CdbParameters(props *OracleSe2CdbInstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_PostgresParameters(props *PostgresInstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_SqlServerEeParameters(props *SqlServerEeInstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_SqlServerExParameters(props *SqlServerExInstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_SqlServerSeParameters(props *SqlServerSeInstanceEngineProps) error {
	return nil
}

func validateDatabaseInstanceEngine_SqlServerWebParameters(props *SqlServerWebInstanceEngineProps) error {
	return nil
}

