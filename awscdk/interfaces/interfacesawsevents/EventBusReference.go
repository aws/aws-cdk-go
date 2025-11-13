package interfacesawsevents


// A reference to a EventBus resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBusReference := &EventBusReference{
//   	EventBusArn: jsii.String("eventBusArn"),
//   	EventBusName: jsii.String("eventBusName"),
//   }
//
type EventBusReference struct {
	// The ARN of the EventBus resource.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
	// The Name of the EventBus resource.
	EventBusName *string `field:"required" json:"eventBusName" yaml:"eventBusName"`
}

