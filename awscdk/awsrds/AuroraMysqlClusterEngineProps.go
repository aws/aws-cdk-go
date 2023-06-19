package awsrds


// Creation properties of the Aurora MySQL database cluster engine.
//
// Used in {@link DatabaseClusterEngine.auroraMysql}.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_2_08_1(),
//   	}),
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("clusteradmin")),
//   	 // Optional - will default to 'admin' username and generated password
//   	InstanceProps: &InstanceProps{
//   		// optional , defaults to t3.medium
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_SMALL),
//   		VpcSubnets: &SubnetSelection{
//   			SubnetType: ec2.SubnetType_PRIVATE_WITH_NAT,
//   		},
//   		Vpc: *Vpc,
//   	},
//   })
//
// Experimental.
type AuroraMysqlClusterEngineProps struct {
	// The version of the Aurora MySQL cluster engine.
	// Experimental.
	Version AuroraMysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

