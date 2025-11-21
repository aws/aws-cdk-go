package awsgreengrass


// A device definition version contains a list of [devices](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-device.html) .
//
// > After you create a device definition version that contains the devices you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an CloudFormation template, `DeviceDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::DeviceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceDefinitionVersionProperty := &DeviceDefinitionVersionProperty{
//   	Devices: []interface{}{
//   		&DeviceProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			Id: jsii.String("id"),
//   			ThingArn: jsii.String("thingArn"),
//
//   			// the properties below are optional
//   			SyncShadow: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-devicedefinitionversion.html
//
type CfnDeviceDefinition_DeviceDefinitionVersionProperty struct {
	// The devices in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-devicedefinitionversion.html#cfn-greengrass-devicedefinition-devicedefinitionversion-devices
	//
	Devices interface{} `field:"required" json:"devices" yaml:"devices"`
}

