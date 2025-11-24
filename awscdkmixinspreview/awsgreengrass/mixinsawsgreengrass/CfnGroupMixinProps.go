package mixinsawsgreengrass


// Properties for CfnGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnGroupMixinProps := &CfnGroupMixinProps{
//   	InitialVersion: &GroupVersionProperty{
//   		ConnectorDefinitionVersionArn: jsii.String("connectorDefinitionVersionArn"),
//   		CoreDefinitionVersionArn: jsii.String("coreDefinitionVersionArn"),
//   		DeviceDefinitionVersionArn: jsii.String("deviceDefinitionVersionArn"),
//   		FunctionDefinitionVersionArn: jsii.String("functionDefinitionVersionArn"),
//   		LoggerDefinitionVersionArn: jsii.String("loggerDefinitionVersionArn"),
//   		ResourceDefinitionVersionArn: jsii.String("resourceDefinitionVersionArn"),
//   		SubscriptionDefinitionVersionArn: jsii.String("subscriptionDefinitionVersionArn"),
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html
//
type CfnGroupMixinProps struct {
	// The group version to include when the group is created.
	//
	// A group version references the Amazon Resource Name (ARN) of a core definition version, device definition version, subscription definition version, and other version types. The group version must reference a core definition version that contains one core. Other version types are optionally included, depending on your business need.
	//
	// > To associate a group version after the group is created, create an [`AWS::Greengrass::GroupVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html) resource and specify the ID of this group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html#cfn-greengrass-group-initialversion
	//
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// The name of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html#cfn-greengrass-group-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM role attached to the group.
	//
	// This role contains the permissions that Lambda functions and connectors use to interact with other AWS services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html#cfn-greengrass-group-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Application-specific metadata to attach to the group.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html#cfn-greengrass-group-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

