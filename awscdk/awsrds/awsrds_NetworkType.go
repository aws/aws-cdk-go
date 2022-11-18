package awsrds


// The network type of the DB instance.
//
// Example:
//   var vpc vpc
//   // VPC and subnets must have IPv6 CIDR blocks
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	engine: rds.databaseClusterEngine.auroraMysql(&auroraMysqlClusterEngineProps{
//   		version: rds.auroraMysqlEngineVersion_VER_3_02_1(),
//   	}),
//   	instanceProps: &instanceProps{
//   		vpc: vpc,
//   		publiclyAccessible: jsii.Boolean(false),
//   	},
//   	networkType: rds.networkType_DUAL,
//   })
//
type NetworkType string

const (
	// IPv4 only network type.
	NetworkType_IPV4 NetworkType = "IPV4"
	// Dual-stack network type.
	NetworkType_DUAL NetworkType = "DUAL"
)

