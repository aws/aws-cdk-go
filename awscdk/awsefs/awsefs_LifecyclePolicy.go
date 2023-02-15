package awsefs


// EFS Lifecycle Policy, if a file is not accessed for given days, it will move to EFS Infrequent Access.
//
// Example:
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &fileSystemProps{
//   	vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	lifecyclePolicy: efs.lifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to infrequent access (IA) storage by default
//   	performanceMode: efs.performanceMode_GENERAL_PURPOSE,
//   	 // default
//   	outOfInfrequentAccessPolicy: efs.outOfInfrequentAccessPolicy_AFTER_1_ACCESS,
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
)

