package interfacesawsdataexchange


// A reference to a EventAction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventActionReference := &EventActionReference{
//   	EventActionArn: jsii.String("eventActionArn"),
//   }
//
type EventActionReference struct {
	// The Arn of the EventAction resource.
	EventActionArn *string `field:"required" json:"eventActionArn" yaml:"eventActionArn"`
}

