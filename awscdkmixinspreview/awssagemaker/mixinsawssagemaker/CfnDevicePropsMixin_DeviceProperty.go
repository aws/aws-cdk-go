package mixinsawssagemaker


// Information of a particular device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deviceProperty := &DeviceProperty{
//   	Description: jsii.String("description"),
//   	DeviceName: jsii.String("deviceName"),
//   	IotThingName: jsii.String("iotThingName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-device-device.html
//
type CfnDevicePropsMixin_DeviceProperty struct {
	// Description of the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-device-device.html#cfn-sagemaker-device-device-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-device-device.html#cfn-sagemaker-device-device-devicename
	//
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// AWS Internet of Things (IoT) object name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-device-device.html#cfn-sagemaker-device-device-iotthingname
	//
	IotThingName *string `field:"optional" json:"iotThingName" yaml:"iotThingName"`
}

