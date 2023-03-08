package awsiot1click


// Properties for defining a `CfnDevice`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceProps := &cfnDeviceProps{
//   	deviceId: jsii.String("deviceId"),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnDeviceProps struct {
	// The ID of the device, such as `G030PX0312744DWM` .
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// A Boolean value indicating whether the device is enabled ( `true` ) or not ( `false` ).
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

