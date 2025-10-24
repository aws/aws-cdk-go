package awsrds


// Properties for MySQL instance engines.
//
// Used in `DatabaseInstanceEngine.mysql`.
//
// Example:
//   var vpc Vpc
//   var kmsKey Key
//
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_39(),
//   	}),
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R7G, ec2.InstanceSize_LARGE),
//   	Vpc: Vpc,
//   	EnablePerformanceInsights: jsii.Boolean(true),
//   	PerformanceInsightRetention: rds.PerformanceInsightRetention_LONG_TERM,
//   	PerformanceInsightEncryptionKey: kmsKey,
//   })
//
type MySqlInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version MysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

