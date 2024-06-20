package awssecuritylake


// Properties for defining a `CfnSubscriberNotification`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriberNotificationProps := &CfnSubscriberNotificationProps{
//   	SubscriberArn: jsii.String("subscriberArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscribernotification.html
//
type CfnSubscriberNotificationProps struct {
	// The Amazon Resource Name (ARN) of the Security Lake subscriber.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscribernotification.html#cfn-securitylake-subscribernotification-subscriberarn
	//
	SubscriberArn *string `field:"required" json:"subscriberArn" yaml:"subscriberArn"`
}

