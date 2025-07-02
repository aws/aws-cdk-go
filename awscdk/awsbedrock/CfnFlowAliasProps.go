package awsbedrock


// Properties for defining a `CfnFlowAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowAliasProps := &CfnFlowAliasProps{
//   	FlowArn: jsii.String("flowArn"),
//   	Name: jsii.String("name"),
//   	RoutingConfiguration: []interface{}{
//   		&FlowAliasRoutingConfigurationListItemProperty{
//   			FlowVersion: jsii.String("flowVersion"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ConcurrencyConfiguration: &FlowAliasConcurrencyConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		MaxConcurrency: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html
//
type CfnFlowAliasProps struct {
	// The Amazon Resource Name (ARN) of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-flowarn
	//
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of configurations about the versions that the alias maps to.
	//
	// Currently, you can only specify one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-routingconfiguration
	//
	RoutingConfiguration interface{} `field:"required" json:"routingConfiguration" yaml:"routingConfiguration"`
	// The configuration that specifies how nodes in the flow are executed concurrently.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-concurrencyconfiguration
	//
	ConcurrencyConfiguration interface{} `field:"optional" json:"concurrencyConfiguration" yaml:"concurrencyConfiguration"`
	// A description of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

