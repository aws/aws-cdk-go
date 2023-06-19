package awsefs


// EFS Lifecycle Policy, if a file is not accessed for given days, it will move to EFS Infrequent Access.
//
// Example:
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
//   	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	LifecyclePolicy: efs.LifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to infrequent access (IA) storage by default
//   	PerformanceMode: efs.PerformanceMode_GENERAL_PURPOSE,
//   	 // default
//   	OutOfInfrequentAccessPolicy: efs.OutOfInfrequentAccessPolicy_AFTER_1_ACCESS,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-elasticfilesystem-filesystem-lifecyclepolicies
//
// Experimental.
type LifecyclePolicy string

const (
	// After 7 days of not being accessed.
	// Experimental.
	LifecyclePolicy_AFTER_7_DAYS LifecyclePolicy = "AFTER_7_DAYS"
	// After 14 days of not being accessed.
	// Experimental.
	LifecyclePolicy_AFTER_14_DAYS LifecyclePolicy = "AFTER_14_DAYS"
	// After 30 days of not being accessed.
	// Experimental.
	LifecyclePolicy_AFTER_30_DAYS LifecyclePolicy = "AFTER_30_DAYS"
	// After 60 days of not being accessed.
	// Experimental.
	LifecyclePolicy_AFTER_60_DAYS LifecyclePolicy = "AFTER_60_DAYS"
	// After 90 days of not being accessed.
	// Experimental.
	LifecyclePolicy_AFTER_90_DAYS LifecyclePolicy = "AFTER_90_DAYS"
)

