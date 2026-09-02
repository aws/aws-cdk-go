package interfacesawspersonalize


// A reference to a EventTracker resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventTrackerReference := &EventTrackerReference{
//   	EventTrackerArn: jsii.String("eventTrackerArn"),
//   }
//
type EventTrackerReference struct {
	// The EventTrackerArn of the EventTracker resource.
	EventTrackerArn *string `field:"required" json:"eventTrackerArn" yaml:"eventTrackerArn"`
}

