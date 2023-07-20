package awsgreengrass


// Properties for defining a `CfnDeviceDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceDefinitionVersionProps := &CfnDeviceDefinitionVersionProps{
//   	DeviceDefinitionId: jsii.String("deviceDefinitionId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html
//
type CfnDeviceDefinitionVersionProps struct {
	// The ID of the device definition associated with this version.
	//
	// This value is a GUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html#cfn-greengrass-devicedefinitionversion-devicedefinitionid
	//
	DeviceDefinitionId *string `field:"required" json:"deviceDefinitionId" yaml:"deviceDefinitionId"`
	// The devices in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html#cfn-greengrass-devicedefinitionversion-devices
	//
	Devices interface{} `field:"required" json:"devices" yaml:"devices"`
}

