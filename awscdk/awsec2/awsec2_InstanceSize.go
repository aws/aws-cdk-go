package awsec2


// What size of instance to use.
//
// Example:
//   var vpc vpc
//
//   var sourceInstance databaseInstance
//
//   rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("Instance"), &DatabaseInstanceFromSnapshotProps{
//   	SnapshotIdentifier: jsii.String("my-snapshot"),
//   	Engine: rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
//   		Version: rds.PostgresEngineVersion_VER_12_3(),
//   	}),
//   	// optional, defaults to m5.large
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_LARGE),
//   	Vpc: Vpc,
//   })
//   rds.NewDatabaseInstanceReadReplica(this, jsii.String("ReadReplica"), &DatabaseInstanceReadReplicaProps{
//   	SourceDatabaseInstance: sourceInstance,
//   	InstanceType: ec2.InstanceType_*Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_LARGE),
//   	Vpc: Vpc,
//   })
//
// Experimental.
type InstanceSize string

const (
	// Instance size NANO (nano).
	// Experimental.
	InstanceSize_NANO InstanceSize = "NANO"
	// Instance size MICRO (micro).
	// Experimental.
	InstanceSize_MICRO InstanceSize = "MICRO"
	// Instance size SMALL (small).
	// Experimental.
	InstanceSize_SMALL InstanceSize = "SMALL"
	// Instance size MEDIUM (medium).
	// Experimental.
	InstanceSize_MEDIUM InstanceSize = "MEDIUM"
	// Instance size LARGE (large).
	// Experimental.
	InstanceSize_LARGE InstanceSize = "LARGE"
	// Instance size XLARGE (xlarge).
	// Experimental.
	InstanceSize_XLARGE InstanceSize = "XLARGE"
	// Instance size XLARGE2 (2xlarge).
	// Experimental.
	InstanceSize_XLARGE2 InstanceSize = "XLARGE2"
	// Instance size XLARGE3 (3xlarge).
	// Experimental.
	InstanceSize_XLARGE3 InstanceSize = "XLARGE3"
	// Instance size XLARGE4 (4xlarge).
	// Experimental.
	InstanceSize_XLARGE4 InstanceSize = "XLARGE4"
	// Instance size XLARGE6 (6xlarge).
	// Experimental.
	InstanceSize_XLARGE6 InstanceSize = "XLARGE6"
	// Instance size XLARGE8 (8xlarge).
	// Experimental.
	InstanceSize_XLARGE8 InstanceSize = "XLARGE8"
	// Instance size XLARGE9 (9xlarge).
	// Experimental.
	InstanceSize_XLARGE9 InstanceSize = "XLARGE9"
	// Instance size XLARGE10 (10xlarge).
	// Experimental.
	InstanceSize_XLARGE10 InstanceSize = "XLARGE10"
	// Instance size XLARGE12 (12xlarge).
	// Experimental.
	InstanceSize_XLARGE12 InstanceSize = "XLARGE12"
	// Instance size XLARGE16 (16xlarge).
	// Experimental.
	InstanceSize_XLARGE16 InstanceSize = "XLARGE16"
	// Instance size XLARGE18 (18xlarge).
	// Experimental.
	InstanceSize_XLARGE18 InstanceSize = "XLARGE18"
	// Instance size XLARGE24 (24xlarge).
	// Experimental.
	InstanceSize_XLARGE24 InstanceSize = "XLARGE24"
	// Instance size XLARGE32 (32xlarge).
	// Experimental.
	InstanceSize_XLARGE32 InstanceSize = "XLARGE32"
	// Instance size XLARGE48 (48xlarge).
	// Experimental.
	InstanceSize_XLARGE48 InstanceSize = "XLARGE48"
	// Instance size XLARGE56 (56xlarge).
	// Experimental.
	InstanceSize_XLARGE56 InstanceSize = "XLARGE56"
	// Instance size XLARGE56 (112xlarge).
	// Experimental.
	InstanceSize_XLARGE112 InstanceSize = "XLARGE112"
	// Instance size METAL (metal).
	// Experimental.
	InstanceSize_METAL InstanceSize = "METAL"
)

