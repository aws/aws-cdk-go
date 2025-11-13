package interfacesawsneptune


// A reference to a EventSubscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSubscriptionReference := &EventSubscriptionReference{
//   	SubscriptionName: jsii.String("subscriptionName"),
//   }
//
type EventSubscriptionReference struct {
	// The SubscriptionName of the EventSubscription resource.
	SubscriptionName *string `field:"required" json:"subscriptionName" yaml:"subscriptionName"`
}

