package awsgreengrass


// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnGroupProps := &cfnGroupProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &groupVersionProperty{
//   		connectorDefinitionVersionArn: jsii.String("connectorDefinitionVersionArn"),
//   		coreDefinitionVersionArn: jsii.String("coreDefinitionVersionArn"),
//   		deviceDefinitionVersionArn: jsii.String("deviceDefinitionVersionArn"),
//   		functionDefinitionVersionArn: jsii.String("functionDefinitionVersionArn"),
//   		loggerDefinitionVersionArn: jsii.String("loggerDefinitionVersionArn"),
//   		resourceDefinitionVersionArn: jsii.String("resourceDefinitionVersionArn"),
//   		subscriptionDefinitionVersionArn: jsii.String("subscriptionDefinitionVersionArn"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//   	tags: tags,
//   }
//
type CfnGroupProps struct {
	// The name of the group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The group version to include when the group is created.
	//
	// A group version references the Amazon Resource Name (ARN) of a core definition version, device definition version, subscription definition version, and other version types. The group version must reference a core definition version that contains one core. Other version types are optionally included, depending on your business need.
	//
	// > To associate a group version after the group is created, create an [`AWS::Greengrass::GroupVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html) resource and specify the ID of this group.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// The Amazon Resource Name (ARN) of the IAM role attached to the group.
	//
	// This role contains the permissions that Lambda functions and connectors use to interact with other AWS services.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Application-specific metadata to attach to the group.
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

