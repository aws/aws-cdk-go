package awsefs


// EFS Out Of Infrequent Access Policy, if a file is accessed given times, it will move back to primary storage class.
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
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-lifecyclepolicy.html#cfn-efs-filesystem-lifecyclepolicy-transitiontoprimarystorageclass
//
type OutOfInfrequentAccessPolicy string

const (
	// After 1 access.
	OutOfInfrequentAccessPolicy_AFTER_1_ACCESS OutOfInfrequentAccessPolicy = "AFTER_1_ACCESS"
)

