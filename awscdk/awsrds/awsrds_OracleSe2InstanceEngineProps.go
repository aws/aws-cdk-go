package awsrds


// Properties for Oracle Standard Edition 2 instance engines.
//
// Used in {@link DatabaseInstanceEngine.oracleSe2}.
//
// Example:
//   var vpc vpc
//
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
//   	engine: rds.databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
//   		version: rds.oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
//   	}),
//   	// optional, defaults to m5.large
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE3, ec2.instanceSize_SMALL),
//   	credentials: rds.credentials.fromGeneratedSecret(jsii.String("syscdk")),
//   	 // Optional - will default to 'admin' username and generated password
//   	vpc: vpc,
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
//   	},
//   })
//
type OracleSe2InstanceEngineProps struct {
	// The exact version of the engine to use.
	Version OracleEngineVersion `field:"required" json:"version" yaml:"version"`
}

