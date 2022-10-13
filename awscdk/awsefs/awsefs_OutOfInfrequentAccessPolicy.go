package awsefs


// EFS Out Of Infrequent Access Policy, if a file is accessed given times, it will move back to primary storage class.
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
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-lifecyclepolicy.html#cfn-efs-filesystem-lifecyclepolicy-transitiontoprimarystorageclass
//
type OutOfInfrequentAccessPolicy string

const (
	// After 1 access.
	OutOfInfrequentAccessPolicy_AFTER_1_ACCESS OutOfInfrequentAccessPolicy = "AFTER_1_ACCESS"
)

