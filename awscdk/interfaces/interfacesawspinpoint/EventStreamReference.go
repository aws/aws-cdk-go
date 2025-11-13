package interfacesawspinpoint


// A reference to a EventStream resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventStreamReference := &EventStreamReference{
//   	EventStreamId: jsii.String("eventStreamId"),
//   }
//
type EventStreamReference struct {
	// The Id of the EventStream resource.
	EventStreamId *string `field:"required" json:"eventStreamId" yaml:"eventStreamId"`
}

