package awsbedrock


// Target Agent to invoke with Prompt.
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
	// Arn representation of the Agent Alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptagentresource.html#cfn-bedrock-promptversion-promptagentresource-agentidentifier
	//
	AgentIdentifier *string `field:"required" json:"agentIdentifier" yaml:"agentIdentifier"`
}

