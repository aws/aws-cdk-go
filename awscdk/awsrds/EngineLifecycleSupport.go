package awsrds


// Engine lifecycle support for Amazon RDS and Amazon Aurora.
//
// Example:
//   var vpc iVpc
//
//
//   rds.NewDatabaseCluster(this, jsii.String("DatabaseCluster"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_07_0(),
//   	}),
//   	Writer: rds.ClusterInstance_ServerlessV2(jsii.String("writerInstance")),
//   	Vpc: Vpc,
//   	EngineLifecycleSupport: rds.EngineLifecycleSupport_OPEN_SOURCE_RDS_EXTENDED_SUPPORT,
//   })
//
type EngineLifecycleSupport string

const (
	// Using Amazon RDS extended support.
	EngineLifecycleSupport_OPEN_SOURCE_RDS_EXTENDED_SUPPORT EngineLifecycleSupport = "OPEN_SOURCE_RDS_EXTENDED_SUPPORT"
	// Not using Amazon RDS extended support.
	EngineLifecycleSupport_OPEN_SOURCE_RDS_EXTENDED_SUPPORT_DISABLED EngineLifecycleSupport = "OPEN_SOURCE_RDS_EXTENDED_SUPPORT_DISABLED"
)

