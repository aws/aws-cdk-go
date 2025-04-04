package awsbedrock


// Contains specifications for an Amazon Bedrock agent with which to use the prompt.
//
// For more information, see [Create a prompt using Prompt management](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-create.html) and [Automate tasks in your application using conversational agents](https://docs.aws.amazon.com/bedrock/latest/userguide/agents.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptAgentResourceProperty := &PromptAgentResourceProperty{
//   	AgentIdentifier: jsii.String("agentIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptagentresource.html
//
type CfnPromptVersion_PromptAgentResourceProperty struct {
	// The ARN of the agent with which to use the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptagentresource.html#cfn-bedrock-promptversion-promptagentresource-agentidentifier
	//
	AgentIdentifier *string `field:"required" json:"agentIdentifier" yaml:"agentIdentifier"`
}

