package awsrds


// The orchestration of updates of multiple instances.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(map[string]instanceType{
//   		"instanceType": ec2.*instanceType_of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
//   	}),
//   	Readers: Readers,
//   	Rds.ClusterInstance_Provisioned(jsii.String("reader")): ,
//   	InstanceUpdateBehaviour: *Rds.InstanceUpdateBehaviour_ROLLING,
//   	 // Optional - defaults to rds.InstanceUpdateBehaviour.BULK
//   	Vpc: Vpc,
//   })
//
type InstanceUpdateBehaviour string

const (
	// In a bulk update, all instances of the cluster are updated at the same time.
	//
	// This results in a faster update procedure.
	// During the update, however, all instances might be unavailable at the same time and thus a downtime might occur.
	InstanceUpdateBehaviour_BULK InstanceUpdateBehaviour = "BULK"
	// In a rolling update, one instance after another is updated.
	//
	// This results in at most one instance being unavailable during the update.
	// If your cluster consists of more than 1 instance, the downtime periods are limited to the time a primary switch needs.
	InstanceUpdateBehaviour_ROLLING InstanceUpdateBehaviour = "ROLLING"
)

