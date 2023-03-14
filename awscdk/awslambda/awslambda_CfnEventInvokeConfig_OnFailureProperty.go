package awslambda


// A destination for events that failed processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onFailureProperty := &onFailureProperty{
//   	destination: jsii.String("destination"),
//   }
//
type CfnEventInvokeConfig_OnFailureProperty struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

