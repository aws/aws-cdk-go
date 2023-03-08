package awsgreengrass


// Properties for defining a `CfnDeviceDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnDeviceDefinitionProps := &cfnDeviceDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &deviceDefinitionVersionProperty{
//   		devices: []interface{}{
//   			&deviceProperty{
//   				certificateArn: jsii.String("certificateArn"),
//   				id: jsii.String("id"),
//   				thingArn: jsii.String("thingArn"),
//
//   				// the properties below are optional
//   				syncShadow: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnDeviceDefinitionProps struct {
	// The name of the device definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The device definition version to include when the device definition is created.
	//
	// A device definition version contains a list of [`device`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-device.html) property types.
	//
	// > To associate a device definition version after the device definition is created, create an [`AWS::Greengrass::DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html) resource and specify the ID of this device definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the device definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

