package mixinsawsbedrock


// Properties for CfnAgentAliasPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentAliasMixinProps := &CfnAgentAliasMixinProps{
//   	AgentAliasName: jsii.String("agentAliasName"),
//   	AgentId: jsii.String("agentId"),
//   	Description: jsii.String("description"),
//   	RoutingConfiguration: []interface{}{
//   		&AgentAliasRoutingConfigurationListItemProperty{
//   			AgentVersion: jsii.String("agentVersion"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agentalias.html
//
type CfnAgentAliasMixinProps struct {
	// The name of the alias of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agentalias.html#cfn-bedrock-agentalias-agentaliasname
	//
	AgentAliasName *string `field:"optional" json:"agentAliasName" yaml:"agentAliasName"`
	// The unique identifier of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agentalias.html#cfn-bedrock-agentalias-agentid
	//
	AgentId *string `field:"optional" json:"agentId" yaml:"agentId"`
	// The description of the alias of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agentalias.html#cfn-bedrock-agentalias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Contains details about the routing configuration of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agentalias.html#cfn-bedrock-agentalias-routingconfiguration
	//
	RoutingConfiguration interface{} `field:"optional" json:"routingConfiguration" yaml:"routingConfiguration"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agentalias.html#cfn-bedrock-agentalias-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

