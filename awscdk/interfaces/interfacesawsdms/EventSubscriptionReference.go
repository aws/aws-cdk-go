package interfacesawsdms


// A reference to a EventSubscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSubscriptionReference := &EventSubscriptionReference{
//   	EventSubscriptionId: jsii.String("eventSubscriptionId"),
//   }
//
type EventSubscriptionReference struct {
	// The Id of the EventSubscription resource.
	EventSubscriptionId *string `field:"required" json:"eventSubscriptionId" yaml:"eventSubscriptionId"`
}

