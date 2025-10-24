package awsrds


// Engine lifecycle support for Amazon RDS and Amazon Aurora.
//
// Example:
//   var vpc IVpc
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
//   rds.NewDatabaseInstance(this, jsii.String("DatabaseInstance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_39(),
//   	}),
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R7G, ec2.InstanceSize_LARGE),
//   	Vpc: Vpc,
//   	EngineLifecycleSupport: rds.EngineLifecycleSupport_OPEN_SOURCE_RDS_EXTENDED_SUPPORT_DISABLED,
//   })
//
type EngineLifecycleSupport string

const (
	// Using Amazon RDS extended support.
	EngineLifecycleSupport_OPEN_SOURCE_RDS_EXTENDED_SUPPORT EngineLifecycleSupport = "OPEN_SOURCE_RDS_EXTENDED_SUPPORT"
	// Not using Amazon RDS extended support.
	EngineLifecycleSupport_OPEN_SOURCE_RDS_EXTENDED_SUPPORT_DISABLED EngineLifecycleSupport = "OPEN_SOURCE_RDS_EXTENDED_SUPPORT_DISABLED"
)

