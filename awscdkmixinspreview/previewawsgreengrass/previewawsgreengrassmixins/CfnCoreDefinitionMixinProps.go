package previewawsgreengrassmixins


// Properties for CfnCoreDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnCoreDefinitionMixinProps := &CfnCoreDefinitionMixinProps{
//   	InitialVersion: &CoreDefinitionVersionProperty{
//   		Cores: []interface{}{
//   			&CoreProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html
//
type CfnCoreDefinitionMixinProps struct {
	// The core definition version to include when the core definition is created.
	//
	// Currently, a core definition version can contain only one [`core`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html) .
	//
	// > To associate a core definition version after the core definition is created, create an [`AWS::Greengrass::CoreDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html) resource and specify the ID of this core definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html#cfn-greengrass-coredefinition-initialversion
	//
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// The name of the core definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html#cfn-greengrass-coredefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Application-specific metadata to attach to the core definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html#cfn-greengrass-coredefinition-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

