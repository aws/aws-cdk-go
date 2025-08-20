package awsgreengrass


// Properties for defining a `CfnConnectorDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var tags interface{}
//
//   cfnConnectorDefinitionProps := &CfnConnectorDefinitionProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	InitialVersion: &ConnectorDefinitionVersionProperty{
//   		Connectors: []interface{}{
//   			&ConnectorProperty{
//   				ConnectorArn: jsii.String("connectorArn"),
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				Parameters: parameters,
//   			},
//   		},
//   	},
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html
//
type CfnConnectorDefinitionProps struct {
	// The name of the connector definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html#cfn-greengrass-connectordefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The connector definition version to include when the connector definition is created.
	//
	// A connector definition version contains a list of [`connector`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinition-connector.html) property types.
	//
	// > To associate a connector definition version after the connector definition is created, create an [`AWS::Greengrass::ConnectorDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html) resource and specify the ID of this connector definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html#cfn-greengrass-connectordefinition-initialversion
	//
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the connector definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html#cfn-greengrass-connectordefinition-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

