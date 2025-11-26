package previewawsdeadlinemixins


// Identifies the user for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jobRunAsUserProperty := &JobRunAsUserProperty{
//   	Posix: &PosixUserProperty{
//   		Group: jsii.String("group"),
//   		User: jsii.String("user"),
//   	},
//   	RunAs: jsii.String("runAs"),
//   	Windows: &WindowsUserProperty{
//   		PasswordArn: jsii.String("passwordArn"),
//   		User: jsii.String("user"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-jobrunasuser.html
//
type CfnQueuePropsMixin_JobRunAsUserProperty struct {
	// The user and group that the jobs in the queue run as.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-jobrunasuser.html#cfn-deadline-queue-jobrunasuser-posix
	//
	Posix interface{} `field:"optional" json:"posix" yaml:"posix"`
	// Specifies whether the job should run using the queue's system user or if the job should run using the worker agent system user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-jobrunasuser.html#cfn-deadline-queue-jobrunasuser-runas
	//
	RunAs *string `field:"optional" json:"runAs" yaml:"runAs"`
	// Identifies a Microsoft Windows user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-jobrunasuser.html#cfn-deadline-queue-jobrunasuser-windows
	//
	Windows interface{} `field:"optional" json:"windows" yaml:"windows"`
}

