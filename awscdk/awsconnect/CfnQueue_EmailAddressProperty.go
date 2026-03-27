package awsconnect


// An email address configuration for the queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailAddressProperty := &EmailAddressProperty{
//   	EmailAddressArn: jsii.String("emailAddressArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-emailaddress.html
//
type CfnQueue_EmailAddressProperty struct {
	// The Amazon Resource Name (ARN) of the email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-emailaddress.html#cfn-connect-queue-emailaddress-emailaddressarn
	//
	EmailAddressArn *string `field:"required" json:"emailAddressArn" yaml:"emailAddressArn"`
}

