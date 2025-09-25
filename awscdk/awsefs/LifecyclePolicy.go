package awsefs


// EFS Lifecycle Policy, if a file is not accessed for given days, it will move to EFS Infrequent Access or Archive storage.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-elasticfilesystem-filesystem-lifecyclepolicies
//
type LifecyclePolicy string

const (
	// After 1 day of not being accessed.
	LifecyclePolicy_AFTER_1_DAY LifecyclePolicy = "AFTER_1_DAY"
	// After 7 days of not being accessed.
	LifecyclePolicy_AFTER_7_DAYS LifecyclePolicy = "AFTER_7_DAYS"
	// After 14 days of not being accessed.
	LifecyclePolicy_AFTER_14_DAYS LifecyclePolicy = "AFTER_14_DAYS"
	// After 30 days of not being accessed.
	LifecyclePolicy_AFTER_30_DAYS LifecyclePolicy = "AFTER_30_DAYS"
	// After 60 days of not being accessed.
	LifecyclePolicy_AFTER_60_DAYS LifecyclePolicy = "AFTER_60_DAYS"
	// After 90 days of not being accessed.
	LifecyclePolicy_AFTER_90_DAYS LifecyclePolicy = "AFTER_90_DAYS"
	// After 180 days of not being accessed.
	LifecyclePolicy_AFTER_180_DAYS LifecyclePolicy = "AFTER_180_DAYS"
	// After 270 days of not being accessed.
	LifecyclePolicy_AFTER_270_DAYS LifecyclePolicy = "AFTER_270_DAYS"
	// After 365 days of not being accessed.
	LifecyclePolicy_AFTER_365_DAYS LifecyclePolicy = "AFTER_365_DAYS"
)

