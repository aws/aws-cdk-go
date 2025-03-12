package awsdeadline


// The POSIX user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   posixUserProperty := &PosixUserProperty{
//   	Group: jsii.String("group"),
//   	User: jsii.String("user"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-posixuser.html
//
type CfnQueue_PosixUserProperty struct {
	// The name of the POSIX user's group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-posixuser.html#cfn-deadline-queue-posixuser-group
	//
	Group *string `field:"required" json:"group" yaml:"group"`
	// The name of the POSIX user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-posixuser.html#cfn-deadline-queue-posixuser-user
	//
	User *string `field:"required" json:"user" yaml:"user"`
}

