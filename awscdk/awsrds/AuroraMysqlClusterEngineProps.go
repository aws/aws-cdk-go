package awsrds


// Creation properties of the Aurora MySQL database cluster engine.
//
// Used in `DatabaseClusterEngine.auroraMysql`.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("Instance"), &ProvisionedClusterInstanceProps{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
//   	}),
//   	Readers: []iClusterInstance{
//   		rds.ClusterInstance_*Provisioned(jsii.String("reader")),
//   	},
//   	InstanceUpdateBehaviour: rds.InstanceUpdateBehaviour_ROLLING,
//   	 // Optional - defaults to rds.InstanceUpdateBehaviour.BULK
//   	Vpc: Vpc,
//   })
//
type AuroraMysqlClusterEngineProps struct {
	// The version of the Aurora MySQL cluster engine.
	Version AuroraMysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

