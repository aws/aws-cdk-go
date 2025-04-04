package awsdeadline


// The Windows user details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   windowsUserProperty := &WindowsUserProperty{
//   	PasswordArn: jsii.String("passwordArn"),
//   	User: jsii.String("user"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-windowsuser.html
//
type CfnQueue_WindowsUserProperty struct {
	// The password ARN for the Windows user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-windowsuser.html#cfn-deadline-queue-windowsuser-passwordarn
	//
	PasswordArn *string `field:"required" json:"passwordArn" yaml:"passwordArn"`
	// The user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-windowsuser.html#cfn-deadline-queue-windowsuser-user
	//
	User *string `field:"required" json:"user" yaml:"user"`
}

