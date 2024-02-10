package awsefs


// The status of the file system's replication overwrite protection.
//
// Example:
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
//   	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	LifecyclePolicy: efs.LifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to infrequent access (IA) storage by default
//   	PerformanceMode: efs.PerformanceMode_GENERAL_PURPOSE,
//   	 // default
//   	OutOfInfrequentAccessPolicy: efs.OutOfInfrequentAccessPolicy_AFTER_1_ACCESS,
//   	 // files are not transitioned back from (infrequent access) IA to primary storage by default
//   	TransitionToArchivePolicy: efs.LifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to Archive by default
//   	ReplicationOverwriteProtection: efs.ReplicationOverwriteProtection_ENABLED,
//   })
//
// See: https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemprotection.html
//
type ReplicationOverwriteProtection string

const (
	// Enable the filesystem's replication overwrite protection.
	ReplicationOverwriteProtection_ENABLED ReplicationOverwriteProtection = "ENABLED"
	// Disable the filesystem's replication overwrite protection.
	ReplicationOverwriteProtection_DISABLED ReplicationOverwriteProtection = "DISABLED"
)

