package awsbedrock


// Contains details about the history of the alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentAliasHistoryEventProperty := &AgentAliasHistoryEventProperty{
//   	EndDate: jsii.String("endDate"),
//   	RoutingConfiguration: []interface{}{
//   		&AgentAliasRoutingConfigurationListItemProperty{
//   			AgentVersion: jsii.String("agentVersion"),
//   		},
//   	},
//   	StartDate: jsii.String("startDate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agentalias-agentaliashistoryevent.html
//
type CfnAgentAlias_AgentAliasHistoryEventProperty struct {
	// The date that the alias stopped being associated to the version in the `routingConfiguration` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agentalias-agentaliashistoryevent.html#cfn-bedrock-agentalias-agentaliashistoryevent-enddate
	//
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// Contains details about the version of the agent with which the alias is associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agentalias-agentaliashistoryevent.html#cfn-bedrock-agentalias-agentaliashistoryevent-routingconfiguration
	//
	RoutingConfiguration interface{} `field:"optional" json:"routingConfiguration" yaml:"routingConfiguration"`
	// The date that the alias began being associated to the version in the `routingConfiguration` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agentalias-agentaliashistoryevent.html#cfn-bedrock-agentalias-agentaliashistoryevent-startdate
	//
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
}

