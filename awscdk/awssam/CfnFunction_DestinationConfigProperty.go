package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigProperty := &DestinationConfigProperty{
//   	OnFailure: &DestinationProperty{
//   		Destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		Type: jsii.String("type"),
//   	},
//   }
//
type CfnFunction_DestinationConfigProperty struct {
	// `CfnFunction.DestinationConfigProperty.OnFailure`.
	OnFailure interface{} `field:"required" json:"onFailure" yaml:"onFailure"`
}

