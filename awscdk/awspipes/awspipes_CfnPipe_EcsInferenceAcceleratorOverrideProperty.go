package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsInferenceAcceleratorOverrideProperty := &ecsInferenceAcceleratorOverrideProperty{
//   	deviceName: jsii.String("deviceName"),
//   	deviceType: jsii.String("deviceType"),
//   }
//
type CfnPipe_EcsInferenceAcceleratorOverrideProperty struct {
	// `CfnPipe.EcsInferenceAcceleratorOverrideProperty.DeviceName`.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// `CfnPipe.EcsInferenceAcceleratorOverrideProperty.DeviceType`.
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

