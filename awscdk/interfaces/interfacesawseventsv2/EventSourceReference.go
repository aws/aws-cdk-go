package interfacesawseventsv2


// A reference to a EventSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceReference := &EventSourceReference{
//   	EventSourceArn: jsii.String("eventSourceArn"),
//   }
//
type EventSourceReference struct {
	// The EventSourceArn of the EventSource resource.
	EventSourceArn *string `field:"required" json:"eventSourceArn" yaml:"eventSourceArn"`
}

