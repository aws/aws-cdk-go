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
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("clusteradmin")),
//   	 // Optional - will default to 'admin' username and generated password
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   		PubliclyAccessible: jsii.Boolean(false),
//   	}),
//   	Readers: []iClusterInstance{
//   		rds.ClusterInstance_*Provisioned(jsii.String("reader1"), &ProvisionedClusterInstanceProps{
//   			PromotionTier: jsii.Number(1),
//   		}),
//   		rds.ClusterInstance_ServerlessV2(jsii.String("reader2")),
//   	},
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   	Vpc: Vpc,
//   })
//
type AuroraMysqlClusterEngineProps struct {
	// The version of the Aurora MySQL cluster engine.
	Version AuroraMysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

