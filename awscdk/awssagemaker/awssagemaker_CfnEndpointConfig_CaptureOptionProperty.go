package awssagemaker


// Specifies whether the endpoint captures input data or output data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captureOptionProperty := &captureOptionProperty{
//   	captureMode: jsii.String("captureMode"),
//   }
//
type CfnEndpointConfig_CaptureOptionProperty struct {
	// Specifies whether the endpoint captures input data or output data.
	CaptureMode *string `field:"required" json:"captureMode" yaml:"captureMode"`
}

