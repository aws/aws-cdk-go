package awsrds


// Creation properties of the Aurora MySQL database cluster engine.
//
// Used in {@link DatabaseClusterEngine.auroraMysql}.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	engine: rds.databaseClusterEngine.auroraMysql(&auroraMysqlClusterEngineProps{
//   		version: rds.auroraMysqlEngineVersion_VER_2_08_1(),
//   	}),
//   	credentials: rds.credentials.fromGeneratedSecret(jsii.String("clusteradmin")),
//   	 // Optional - will default to 'admin' username and generated password
//   	instanceProps: &instanceProps{
//   		// optional , defaults to t3.medium
//   		instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_SMALL),
//   		vpcSubnets: &subnetSelection{
//   			subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
//   		},
//   		vpc: vpc,
//   	},
//   })
//
type AuroraMysqlClusterEngineProps struct {
	// The version of the Aurora MySQL cluster engine.
	Version AuroraMysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

