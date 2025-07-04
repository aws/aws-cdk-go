package awseventstargets


// Override inference accelerators for the task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceAcceleratorOverride := &InferenceAcceleratorOverride{
//   	DeviceName: jsii.String("deviceName"),
//   	DeviceType: jsii.String("deviceType"),
//   }
//
type InferenceAcceleratorOverride struct {
	// The Elastic Inference accelerator device name to override for the task.
	//
	// This parameter must match a `deviceName` specified in the task definition.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// The Elastic Inference accelerator type to use.
	DeviceType *string `field:"required" json:"deviceType" yaml:"deviceType"`
}

