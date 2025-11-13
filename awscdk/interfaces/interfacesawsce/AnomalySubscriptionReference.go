package interfacesawsce


// A reference to a AnomalySubscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anomalySubscriptionReference := &AnomalySubscriptionReference{
//   	SubscriptionArn: jsii.String("subscriptionArn"),
//   }
//
type AnomalySubscriptionReference struct {
	// The SubscriptionArn of the AnomalySubscription resource.
	SubscriptionArn *string `field:"required" json:"subscriptionArn" yaml:"subscriptionArn"`
}

