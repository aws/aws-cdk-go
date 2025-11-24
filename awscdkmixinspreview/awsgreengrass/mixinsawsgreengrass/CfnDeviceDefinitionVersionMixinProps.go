package mixinsawsgreengrass


// Properties for CfnDeviceDefinitionVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeviceDefinitionVersionMixinProps := &CfnDeviceDefinitionVersionMixinProps{
//   	DeviceDefinitionId: jsii.String("deviceDefinitionId"),
//   	Devices: []interface{}{
//   		&DeviceProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			Id: jsii.String("id"),
//   			SyncShadow: jsii.Boolean(false),
//   			ThingArn: jsii.String("thingArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html
//
type CfnDeviceDefinitionVersionMixinProps struct {
	// The ID of the device definition associated with this version.
	//
	// This value is a GUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html#cfn-greengrass-devicedefinitionversion-devicedefinitionid
	//
	DeviceDefinitionId *string `field:"optional" json:"deviceDefinitionId" yaml:"deviceDefinitionId"`
	// The devices in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html#cfn-greengrass-devicedefinitionversion-devices
	//
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
}

