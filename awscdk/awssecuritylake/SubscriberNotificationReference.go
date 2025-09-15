package awssecuritylake


// A reference to a SubscriberNotification resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriberNotificationReference := &SubscriberNotificationReference{
//   	SubscriberArn: jsii.String("subscriberArn"),
//   }
//
type SubscriberNotificationReference struct {
	// The SubscriberArn of the SubscriberNotification resource.
	SubscriberArn *string `field:"required" json:"subscriberArn" yaml:"subscriberArn"`
}

