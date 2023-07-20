package awsiot1click


// Properties for defining a `CfnDevice`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceProps := &CfnDeviceProps{
//   	DeviceId: jsii.String("deviceId"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html
//
type CfnDeviceProps struct {
	// The ID of the device, such as `G030PX0312744DWM` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html#cfn-iot1click-device-deviceid
	//
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// A Boolean value indicating whether the device is enabled ( `true` ) or not ( `false` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html#cfn-iot1click-device-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

