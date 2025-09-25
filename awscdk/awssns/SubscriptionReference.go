package awssns


// A reference to a Subscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionReference := &SubscriptionReference{
//   	SubscriptionArn: jsii.String("subscriptionArn"),
//   }
//
type SubscriptionReference struct {
	// The Arn of the Subscription resource.
	SubscriptionArn *string `field:"required" json:"subscriptionArn" yaml:"subscriptionArn"`
}

