package awsevents


// A reference to a EventBusPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBusPolicyReference := &EventBusPolicyReference{
//   	EventBusPolicyId: jsii.String("eventBusPolicyId"),
//   }
//
type EventBusPolicyReference struct {
	// The Id of the EventBusPolicy resource.
	EventBusPolicyId *string `field:"required" json:"eventBusPolicyId" yaml:"eventBusPolicyId"`
}

