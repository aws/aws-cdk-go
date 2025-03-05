package awsbedrock


// Defines an agent node in your flow.
//
// You specify the agent to invoke at this point in the flow. For more information, see [Node types in Amazon Bedrock works](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-nodes.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentFlowNodeConfigurationProperty := &AgentFlowNodeConfigurationProperty{
//   	AgentAliasArn: jsii.String("agentAliasArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-agentflownodeconfiguration.html
//
type CfnFlowVersion_AgentFlowNodeConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the alias of the agent to invoke.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-agentflownodeconfiguration.html#cfn-bedrock-flowversion-agentflownodeconfiguration-agentaliasarn
	//
	AgentAliasArn *string `field:"required" json:"agentAliasArn" yaml:"agentAliasArn"`
}

