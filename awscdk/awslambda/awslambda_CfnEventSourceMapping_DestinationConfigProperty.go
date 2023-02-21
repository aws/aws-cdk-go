package awslambda


// A configuration object that specifies the destination of an event after Lambda processes it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigProperty := &destinationConfigProperty{
//   	onFailure: &onFailureProperty{
//   		destination: jsii.String("destination"),
//   	},
//   }
//
type CfnEventSourceMapping_DestinationConfigProperty struct {
	// The destination configuration for failed invocations.
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
}

