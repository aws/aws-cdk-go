package mixinsawsbedrock


// Properties for CfnFlowAliasPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowAliasMixinProps := &CfnFlowAliasMixinProps{
//   	ConcurrencyConfiguration: &FlowAliasConcurrencyConfigurationProperty{
//   		MaxConcurrency: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   	Description: jsii.String("description"),
//   	FlowArn: jsii.String("flowArn"),
//   	Name: jsii.String("name"),
//   	RoutingConfiguration: []interface{}{
//   		&FlowAliasRoutingConfigurationListItemProperty{
//   			FlowVersion: jsii.String("flowVersion"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html
//
type CfnFlowAliasMixinProps struct {
	// The configuration that specifies how nodes in the flow are executed concurrently.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-concurrencyconfiguration
	//
	ConcurrencyConfiguration interface{} `field:"optional" json:"concurrencyConfiguration" yaml:"concurrencyConfiguration"`
	// A description of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-flowarn
	//
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The name of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of configurations about the versions that the alias maps to.
	//
	// Currently, you can only specify one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-routingconfiguration
	//
	RoutingConfiguration interface{} `field:"optional" json:"routingConfiguration" yaml:"routingConfiguration"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowalias.html#cfn-bedrock-flowalias-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

