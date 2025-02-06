package awsbedrock


// Contains specifications for a generative AI resource with which to use the prompt.
//
// For more information, see [Create a prompt using Prompt management](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-create.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptGenAiResourceProperty := &PromptGenAiResourceProperty{
//   	Agent: &PromptAgentResourceProperty{
//   		AgentIdentifier: jsii.String("agentIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptgenairesource.html
//
type CfnPromptVersion_PromptGenAiResourceProperty struct {
	// Specifies an Amazon Bedrock agent with which to use the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptgenairesource.html#cfn-bedrock-promptversion-promptgenairesource-agent
	//
	Agent interface{} `field:"required" json:"agent" yaml:"agent"`
}

