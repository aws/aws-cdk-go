package awsgreengrass


// Properties for defining a `CfnCoreDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnCoreDefinitionProps := &cfnCoreDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &coreDefinitionVersionProperty{
//   		cores: []interface{}{
//   			&coreProperty{
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
type CfnCoreDefinitionProps struct {
	// The name of the core definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The core definition version to include when the core definition is created.
	//
	// Currently, a core definition version can contain only one [`core`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html) .
	//
	// > To associate a core definition version after the core definition is created, create an [`AWS::Greengrass::CoreDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html) resource and specify the ID of this core definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the core definition.
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

