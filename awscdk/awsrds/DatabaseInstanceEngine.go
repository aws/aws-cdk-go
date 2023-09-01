package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A database instance engine.
//
// Provides mapping to DatabaseEngine used for
// secret rotation.
//
// Example:
//   var vpc vpc
//
//
//   iopsInstance := rds.NewDatabaseInstance(this, jsii.String("IopsInstance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_30(),
//   	}),
//   	Vpc: Vpc,
//   	StorageType: rds.StorageType_IO1,
//   	Iops: jsii.Number(5000),
//   })
//
//   gp3Instance := rds.NewDatabaseInstance(this, jsii.String("Gp3Instance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_*Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_30(),
//   	}),
//   	Vpc: Vpc,
//   	AllocatedStorage: jsii.Number(500),
//   	StorageType: rds.StorageType_GP3,
//   	StorageThroughput: jsii.Number(500),
//   })
//
type DatabaseInstanceEngine interface {
}

// The jsii proxy struct for DatabaseInstanceEngine
type jsiiProxy_DatabaseInstanceEngine struct {
	_ byte // padding
}

func NewDatabaseInstanceEngine() DatabaseInstanceEngine {
	_init_.Initialize()

	j := jsiiProxy_DatabaseInstanceEngine{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDatabaseInstanceEngine_Override(d DatabaseInstanceEngine) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		nil, // no parameters
		d,
	)
}

// Creates a new MariaDB instance engine.
func DatabaseInstanceEngine_MariaDb(props *MariaDbInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_MariaDbParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"mariaDb",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new MySQL instance engine.
func DatabaseInstanceEngine_Mysql(props *MySqlInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_MysqlParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"mysql",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Enterprise Edition instance engine.
func DatabaseInstanceEngine_OracleEe(props *OracleEeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_OracleEeParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"oracleEe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Enterprise Edition (CDB) instance engine.
func DatabaseInstanceEngine_OracleEeCdb(props *OracleEeCdbInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_OracleEeCdbParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"oracleEeCdb",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Standard Edition 2 instance engine.
func DatabaseInstanceEngine_OracleSe2(props *OracleSe2InstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_OracleSe2Parameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"oracleSe2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Standard Edition 2 (CDB) instance engine.
func DatabaseInstanceEngine_OracleSe2Cdb(props *OracleSe2CdbInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_OracleSe2CdbParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"oracleSe2Cdb",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new PostgreSQL instance engine.
func DatabaseInstanceEngine_Postgres(props *PostgresInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_PostgresParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"postgres",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Enterprise Edition instance engine.
func DatabaseInstanceEngine_SqlServerEe(props *SqlServerEeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_SqlServerEeParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"sqlServerEe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Express Edition instance engine.
func DatabaseInstanceEngine_SqlServerEx(props *SqlServerExInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_SqlServerExParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"sqlServerEx",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Standard Edition instance engine.
func DatabaseInstanceEngine_SqlServerSe(props *SqlServerSeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_SqlServerSeParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"sqlServerSe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Web Edition instance engine.
func DatabaseInstanceEngine_SqlServerWeb(props *SqlServerWebInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_SqlServerWebParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"sqlServerWeb",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func DatabaseInstanceEngine_MARIADB() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"MARIADB",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_MYSQL() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"MYSQL",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_EE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"ORACLE_EE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_EE_CDB() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"ORACLE_EE_CDB",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_SE2() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"ORACLE_SE2",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_SE2_CDB() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"ORACLE_SE2_CDB",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_POSTGRES() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"POSTGRES",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_EE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_EE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_EX() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_EX",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_SE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_SE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_WEB() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_WEB",
		&returns,
	)
	return returns
}

