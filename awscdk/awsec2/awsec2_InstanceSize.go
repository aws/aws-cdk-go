package awsec2


// What size of instance to use.
//
// Example:
//   var vpc vpc
//
//   var sourceInstance databaseInstance
//
//   rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("Instance"), &databaseInstanceFromSnapshotProps{
//   	snapshotIdentifier: jsii.String("my-snapshot"),
//   	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   		version: rds.postgresEngineVersion_VER_12_3(),
//   	}),
//   	// optional, defaults to m5.large
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_LARGE),
//   	vpc: vpc,
//   })
//   rds.NewDatabaseInstanceReadReplica(this, jsii.String("ReadReplica"), &databaseInstanceReadReplicaProps{
//   	sourceDatabaseInstance: sourceInstance,
//   	instanceType: ec2.*instanceType.of(ec2.*instanceClass_BURSTABLE2, ec2.*instanceSize_LARGE),
//   	vpc: vpc,
//   })
//
type InstanceSize string

const (
	// Instance size NANO (nano).
	InstanceSize_NANO InstanceSize = "NANO"
	// Instance size MICRO (micro).
	InstanceSize_MICRO InstanceSize = "MICRO"
	// Instance size SMALL (small).
	InstanceSize_SMALL InstanceSize = "SMALL"
	// Instance size MEDIUM (medium).
	InstanceSize_MEDIUM InstanceSize = "MEDIUM"
	// Instance size LARGE (large).
	InstanceSize_LARGE InstanceSize = "LARGE"
	// Instance size XLARGE (xlarge).
	InstanceSize_XLARGE InstanceSize = "XLARGE"
	// Instance size XLARGE2 (2xlarge).
	InstanceSize_XLARGE2 InstanceSize = "XLARGE2"
	// Instance size XLARGE3 (3xlarge).
	InstanceSize_XLARGE3 InstanceSize = "XLARGE3"
	// Instance size XLARGE4 (4xlarge).
	InstanceSize_XLARGE4 InstanceSize = "XLARGE4"
	// Instance size XLARGE6 (6xlarge).
	InstanceSize_XLARGE6 InstanceSize = "XLARGE6"
	// Instance size XLARGE8 (8xlarge).
	InstanceSize_XLARGE8 InstanceSize = "XLARGE8"
	// Instance size XLARGE9 (9xlarge).
	InstanceSize_XLARGE9 InstanceSize = "XLARGE9"
	// Instance size XLARGE10 (10xlarge).
	InstanceSize_XLARGE10 InstanceSize = "XLARGE10"
	// Instance size XLARGE12 (12xlarge).
	InstanceSize_XLARGE12 InstanceSize = "XLARGE12"
	// Instance size XLARGE16 (16xlarge).
	InstanceSize_XLARGE16 InstanceSize = "XLARGE16"
	// Instance size XLARGE18 (18xlarge).
	InstanceSize_XLARGE18 InstanceSize = "XLARGE18"
	// Instance size XLARGE24 (24xlarge).
	InstanceSize_XLARGE24 InstanceSize = "XLARGE24"
	// Instance size XLARGE32 (32xlarge).
	InstanceSize_XLARGE32 InstanceSize = "XLARGE32"
	// Instance size XLARGE48 (48xlarge).
	InstanceSize_XLARGE48 InstanceSize = "XLARGE48"
	// Instance size XLARGE56 (56xlarge).
	InstanceSize_XLARGE56 InstanceSize = "XLARGE56"
	// Instance size XLARGE56 (112xlarge).
	InstanceSize_XLARGE112 InstanceSize = "XLARGE112"
	// Instance size METAL (metal).
	InstanceSize_METAL InstanceSize = "METAL"
)

