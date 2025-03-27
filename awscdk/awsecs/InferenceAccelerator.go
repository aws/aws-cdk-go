package awsecs


// Elastic Inference Accelerator.
//
// For more information, see [Elastic Inference Basics](https://docs.aws.amazon.com/elastic-inference/latest/developerguide/basics.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceAccelerator := &InferenceAccelerator{
//   	DeviceName: jsii.String("deviceName"),
//   	DeviceType: jsii.String("deviceType"),
//   }
//
type InferenceAccelerator struct {
	// The Elastic Inference accelerator device name.
	// Default: - empty.
	//
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// The Elastic Inference accelerator type to use.
	//
	// The allowed values are: eia2.medium, eia2.large and eia2.xlarge.
	// Default: - empty.
	//
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

