package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
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
//   var sourceInstance databaseInstance
//
//   rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("Instance"), &databaseInstanceFromSnapshotProps{
//   	snapshotIdentifier: jsii.String("my-snapshot"),
//   	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   		version: rds.postgresEngineVersion_VER_12_3(),
//   	}),
//   	// optional, defaults to m5.large
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_LARGE),
//   	vpc: vpc,
//   })
//   rds.NewDatabaseInstanceReadReplica(this, jsii.String("ReadReplica"), &databaseInstanceReadReplicaProps{
//   	sourceDatabaseInstance: sourceInstance,
//   	instanceType: ec2.*instanceType.of(ec2.*instanceClass_BURSTABLE2, ec2.*instanceSize_LARGE),
//   	vpc: vpc,
//   })
//
// Experimental.
type DatabaseInstanceEngine interface {
}

// The jsii proxy struct for DatabaseInstanceEngine
type jsiiProxy_DatabaseInstanceEngine struct {
	_ byte // padding
}

// Experimental.
func NewDatabaseInstanceEngine() DatabaseInstanceEngine {
	_init_.Initialize()

	j := jsiiProxy_DatabaseInstanceEngine{}

	_jsii_.Create(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatabaseInstanceEngine_Override(d DatabaseInstanceEngine) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		nil, // no parameters
		d,
	)
}

// Creates a new MariaDB instance engine.
// Experimental.
func DatabaseInstanceEngine_MariaDb(props *MariaDbInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_MariaDbParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"mariaDb",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new MySQL instance engine.
// Experimental.
func DatabaseInstanceEngine_Mysql(props *MySqlInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_MysqlParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"mysql",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Enterprise Edition instance engine.
// Experimental.
func DatabaseInstanceEngine_OracleEe(props *OracleEeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_OracleEeParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"oracleEe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Standard Edition instance engine.
// Deprecated: instances can no longer be created with this engine. See https://forums.aws.amazon.com/ann.jspa?annID=7341
func DatabaseInstanceEngine_OracleSe(props *OracleSeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_OracleSeParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"oracleSe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Standard Edition 1 instance engine.
// Deprecated: instances can no longer be created with this engine. See https://forums.aws.amazon.com/ann.jspa?annID=7341
func DatabaseInstanceEngine_OracleSe1(props *OracleSe1InstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_OracleSe1Parameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"oracleSe1",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Oracle Standard Edition 1 instance engine.
// Experimental.
func DatabaseInstanceEngine_OracleSe2(props *OracleSe2InstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_OracleSe2Parameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"oracleSe2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new PostgreSQL instance engine.
// Experimental.
func DatabaseInstanceEngine_Postgres(props *PostgresInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_PostgresParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"postgres",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Enterprise Edition instance engine.
// Experimental.
func DatabaseInstanceEngine_SqlServerEe(props *SqlServerEeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_SqlServerEeParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"sqlServerEe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Express Edition instance engine.
// Experimental.
func DatabaseInstanceEngine_SqlServerEx(props *SqlServerExInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_SqlServerExParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"sqlServerEx",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Standard Edition instance engine.
// Experimental.
func DatabaseInstanceEngine_SqlServerSe(props *SqlServerSeInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_SqlServerSeParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"sqlServerSe",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new SQL Server Web Edition instance engine.
// Experimental.
func DatabaseInstanceEngine_SqlServerWeb(props *SqlServerWebInstanceEngineProps) IInstanceEngine {
	_init_.Initialize()

	if err := validateDatabaseInstanceEngine_SqlServerWebParameters(props); err != nil {
		panic(err)
	}
	var returns IInstanceEngine

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseInstanceEngine",
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
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"MARIADB",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_MYSQL() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"MYSQL",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_EE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"ORACLE_EE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_SE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"ORACLE_SE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_SE1() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"ORACLE_SE1",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_ORACLE_SE2() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"ORACLE_SE2",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_POSTGRES() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"POSTGRES",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_EE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_EE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_EX() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_EX",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_SE() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_SE",
		&returns,
	)
	return returns
}

func DatabaseInstanceEngine_SQL_SERVER_WEB() IInstanceEngine {
	_init_.Initialize()
	var returns IInstanceEngine
	_jsii_.StaticGet(
		"monocdk.aws_rds.DatabaseInstanceEngine",
		"SQL_SERVER_WEB",
		&returns,
	)
	return returns
}

