package awslambda


// A destination for events that were processed successfully.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onSuccessProperty := &onSuccessProperty{
//   	destination: jsii.String("destination"),
//   }
//
type CfnEventInvokeConfig_OnSuccessProperty struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

