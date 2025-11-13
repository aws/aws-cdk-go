package interfacesawsfrauddetector


// A reference to a EventType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventTypeReference := &EventTypeReference{
//   	EventTypeArn: jsii.String("eventTypeArn"),
//   }
//
type EventTypeReference struct {
	// The Arn of the EventType resource.
	EventTypeArn *string `field:"required" json:"eventTypeArn" yaml:"eventTypeArn"`
}

