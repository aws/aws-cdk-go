package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Abstract base class for prompt GenAI resource configurations.
//
// This provides a high-level abstraction over the underlying CloudFormation
// GenAI resource properties, offering a more developer-friendly interface
// while maintaining full compatibility with the underlying AWS Bedrock service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var agentAlias agentAlias
//
//   promptGenAiResource := bedrock_alpha.PromptGenAiResource_Agent(&AgentGenAiResourceProps{
//   	AgentAlias: agentAlias,
//   })
//
// Experimental.
type PromptGenAiResource interface {
}

// The jsii proxy struct for PromptGenAiResource
type jsiiProxy_PromptGenAiResource struct {
	_ byte // padding
}

// Experimental.
func NewPromptGenAiResource_Override(p PromptGenAiResource) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.PromptGenAiResource",
		nil, // no parameters
		p,
	)
}

// Creates an agent GenAI resource configuration.
// Experimental.
func PromptGenAiResource_Agent(props *AgentGenAiResourceProps) PromptGenAiResource {
	_init_.Initialize()

	if err := validatePromptGenAiResource_AgentParameters(props); err != nil {
		panic(err)
	}
	var returns PromptGenAiResource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptGenAiResource",
		"agent",
		[]interface{}{props},
		&returns,
	)

	return returns
}

