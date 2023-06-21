package awsrds


// Properties for Oracle Standard Edition 2 instance engines.
//
// Used in `DatabaseInstanceEngine.oracleSe2`.
//
// Example:
//   var vpc vpc
//
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_OracleSe2(&OracleSe2InstanceEngineProps{
//   		Version: rds.OracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
//   	}),
//   	// optional, defaults to m5.large
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("syscdk")),
//   	 // Optional - will default to 'admin' username and generated password
//   	Vpc: Vpc,
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   })
//
type OracleSe2InstanceEngineProps struct {
	// The exact version of the engine to use.
	Version OracleEngineVersion `field:"required" json:"version" yaml:"version"`
}

