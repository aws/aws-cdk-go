package awsgreengrass


// A device is an AWS IoT device (thing) that's added to a Greengrass group.
//
// Greengrass devices can communicate with the Greengrass core in the same group. For more information, see [What Is AWS IoT Greengrass ?](https://docs.aws.amazon.com/greengrass/latest/developerguide/what-is-gg.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Devices` property of the [`DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-devicedefinitionversion.html) property type contains a list of `Device` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceProperty := &deviceProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	id: jsii.String("id"),
//   	thingArn: jsii.String("thingArn"),
//
//   	// the properties below are optional
//   	syncShadow: jsii.Boolean(false),
//   }
//
type CfnDeviceDefinition_DeviceProperty struct {
	// The Amazon Resource Name (ARN) of the device certificate for the device.
	//
	// This X.509 certificate is used to authenticate the device with AWS IoT and AWS IoT Greengrass services.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// A descriptive or arbitrary ID for the device.
	//
	// This value must be unique within the device definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ARN of the device, which is an AWS IoT device (thing).
	ThingArn *string `field:"required" json:"thingArn" yaml:"thingArn"`
	// Indicates whether the device's local shadow is synced with the cloud automatically.
	SyncShadow interface{} `field:"optional" json:"syncShadow" yaml:"syncShadow"`
}

