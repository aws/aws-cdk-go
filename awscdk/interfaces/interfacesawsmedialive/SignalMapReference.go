package interfacesawsmedialive


// A reference to a SignalMap resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signalMapReference := &SignalMapReference{
//   	Identifier: jsii.String("identifier"),
//   	SignalMapArn: jsii.String("signalMapArn"),
//   }
//
type SignalMapReference struct {
	// The Identifier of the SignalMap resource.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The ARN of the SignalMap resource.
	SignalMapArn *string `field:"required" json:"signalMapArn" yaml:"signalMapArn"`
}

