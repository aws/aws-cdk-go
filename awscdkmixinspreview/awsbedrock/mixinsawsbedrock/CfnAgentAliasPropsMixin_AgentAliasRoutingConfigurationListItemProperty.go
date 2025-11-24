package mixinsawsbedrock


// Contains details about the routing configuration of the alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   agentAliasRoutingConfigurationListItemProperty := &AgentAliasRoutingConfigurationListItemProperty{
//   	AgentVersion: jsii.String("agentVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agentalias-agentaliasroutingconfigurationlistitem.html
//
type CfnAgentAliasPropsMixin_AgentAliasRoutingConfigurationListItemProperty struct {
	// The version of the agent with which the alias is associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agentalias-agentaliasroutingconfigurationlistitem.html#cfn-bedrock-agentalias-agentaliasroutingconfigurationlistitem-agentversion
	//
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
}

