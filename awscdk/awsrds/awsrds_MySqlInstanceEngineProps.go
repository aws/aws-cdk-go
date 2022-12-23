package awsrds


// Properties for MySQL instance engines.
//
// Used in {@link DatabaseInstanceEngine.mysql}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//
//
//   iopsInstance := rds.NewDatabaseInstance(this, jsii.String("IopsInstance"), &databaseInstanceProps{
//   	engine: rds.databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
//   		version: mysqlEngineVersion_VER_8_0_30,
//   	}),
//   	vpc: vpc,
//   	storageType: rds.storageType_IO1,
//   	iops: jsii.Number(5000),
//   })
//
//   gp3Instance := rds.NewDatabaseInstance(this, jsii.String("Gp3Instance"), &databaseInstanceProps{
//   	engine: rds.*databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
//   		version: *mysqlEngineVersion_VER_8_0_30,
//   	}),
//   	vpc: vpc,
//   	allocatedStorage: jsii.Number(500),
//   	storageType: rds.*storageType_GP3,
//   	storageThroughput: jsii.Number(500),
//   })
//
type MySqlInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version MysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

