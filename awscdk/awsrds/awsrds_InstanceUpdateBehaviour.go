package awsrds


// The orchestration of updates of multiple instances.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	engine: rds.databaseClusterEngine.auroraMysql(&auroraMysqlClusterEngineProps{
//   		version: rds.auroraMysqlEngineVersion_VER_3_01_0(),
//   	}),
//   	instances: jsii.Number(2),
//   	instanceProps: &instanceProps{
//   		instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE3, ec2.instanceSize_SMALL),
//   		vpc: vpc,
//   	},
//   	instanceUpdateBehaviour: rds.instanceUpdateBehaviour_ROLLING,
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

