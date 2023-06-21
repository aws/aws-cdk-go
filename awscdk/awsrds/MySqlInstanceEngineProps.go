package awsrds


// Properties for MySQL instance engines.
//
// Used in `DatabaseInstanceEngine.mysql`.
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
type MySqlInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version MysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

