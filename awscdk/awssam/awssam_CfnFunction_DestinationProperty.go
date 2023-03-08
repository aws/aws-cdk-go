package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &destinationProperty{
//   	destination: jsii.String("destination"),
//
//   	// the properties below are optional
//   	type: jsii.String("type"),
//   }
//
type CfnFunction_DestinationProperty struct {
	// `CfnFunction.DestinationProperty.Destination`.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// `CfnFunction.DestinationProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

