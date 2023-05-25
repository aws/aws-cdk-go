package awsrds


// The network type of the DB instance.
//
// Example:
//   var vpc vpc
//   // VPC and subnets must have IPv6 CIDR blocks
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_02_1(),
//   	}),
//   	InstanceProps: &InstanceProps{
//   		Vpc: *Vpc,
//   		PubliclyAccessible: jsii.Boolean(false),
//   	},
//   	NetworkType: rds.NetworkType_DUAL,
//   })
//
type NetworkType string

const (
	// IPv4 only network type.
	NetworkType_IPV4 NetworkType = "IPV4"
	// Dual-stack network type.
	NetworkType_DUAL NetworkType = "DUAL"
)

