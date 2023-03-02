package awssagemaker


// Information of a particular device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceProperty := &deviceProperty{
//   	deviceName: jsii.String("deviceName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	iotThingName: jsii.String("iotThingName"),
//   }
//
type CfnDevice_DeviceProperty struct {
	// The name of the device.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Description of the device.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS Internet of Things (IoT) object name.
	IotThingName *string `field:"optional" json:"iotThingName" yaml:"iotThingName"`
}

