package awscustomerprofiles


// A reference to a EventTrigger resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventTriggerReference := &EventTriggerReference{
//   	DomainName: jsii.String("domainName"),
//   	EventTriggerName: jsii.String("eventTriggerName"),
//   }
//
type EventTriggerReference struct {
	// The DomainName of the EventTrigger resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The EventTriggerName of the EventTrigger resource.
	EventTriggerName *string `field:"required" json:"eventTriggerName" yaml:"eventTriggerName"`
}

