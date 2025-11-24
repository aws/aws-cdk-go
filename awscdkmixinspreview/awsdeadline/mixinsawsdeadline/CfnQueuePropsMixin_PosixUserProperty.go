package mixinsawsdeadline


// The POSIX user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   posixUserProperty := &PosixUserProperty{
//   	Group: jsii.String("group"),
//   	User: jsii.String("user"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-posixuser.html
//
type CfnQueuePropsMixin_PosixUserProperty struct {
	// The name of the POSIX user's group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-posixuser.html#cfn-deadline-queue-posixuser-group
	//
	Group *string `field:"optional" json:"group" yaml:"group"`
	// The name of the POSIX user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-posixuser.html#cfn-deadline-queue-posixuser-user
	//
	User *string `field:"optional" json:"user" yaml:"user"`
}

