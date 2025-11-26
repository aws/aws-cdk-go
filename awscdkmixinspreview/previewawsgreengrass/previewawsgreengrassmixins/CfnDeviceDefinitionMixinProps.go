package previewawsgreengrassmixins


// Properties for CfnDeviceDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnDeviceDefinitionMixinProps := &CfnDeviceDefinitionMixinProps{
//   	InitialVersion: &DeviceDefinitionVersionProperty{
//   		Devices: []interface{}{
//   			&DeviceProperty{
//   				CertificateArn: jsii.String("certificateArn"),
//   				Id: jsii.String("id"),
//   				SyncShadow: jsii.Boolean(false),
//   				ThingArn: jsii.String("thingArn"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html
//
type CfnDeviceDefinitionMixinProps struct {
	// The device definition version to include when the device definition is created.
	//
	// A device definition version contains a list of [`device`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-device.html) property types.
	//
	// > To associate a device definition version after the device definition is created, create an [`AWS::Greengrass::DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html) resource and specify the ID of this device definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html#cfn-greengrass-devicedefinition-initialversion
	//
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// The name of the device definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html#cfn-greengrass-devicedefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Application-specific metadata to attach to the device definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html#cfn-greengrass-devicedefinition-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

