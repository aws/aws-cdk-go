package awsappintegrations


// The event integration filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventFilterProperty := &eventFilterProperty{
//   	source: jsii.String("source"),
//   }
//
type CfnEventIntegration_EventFilterProperty struct {
	// The source of the events.
	Source *string `field:"required" json:"source" yaml:"source"`
}

