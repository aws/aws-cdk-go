package awsrds


// Creation properties of the Aurora MySQL database cluster engine.
//
// Used in `DatabaseClusterEngine.auroraMysql`.
//
// Example:
//   var vpc Vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   		CaCertificate: rds.CaCertificate_RDS_CA_RSA2048_G1(),
//   	}),
//   	Readers: []IClusterInstance{
//   		rds.ClusterInstance_ServerlessV2(jsii.String("reader"), &ServerlessV2ClusterInstanceProps{
//   			CaCertificate: rds.CaCertificate_Of(jsii.String("custom-ca")),
//   		}),
//   	},
//   	Vpc: Vpc,
//   })
//
type AuroraMysqlClusterEngineProps struct {
	// The version of the Aurora MySQL cluster engine.
	Version AuroraMysqlEngineVersion `field:"required" json:"version" yaml:"version"`
}

