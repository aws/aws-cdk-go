package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2"
)

// Configuration for overriding model and prompt template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var bedrockInvokable IBedrockInvokable
//
//   overrideConfig := &OverrideConfig{
//   	AppendToPrompt: jsii.String("appendToPrompt"),
//   	Model: bedrockInvokable,
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type OverrideConfig struct {
	// The prompt that will be appended to the system prompt to define the model's memory consolidation/extraction strategy.
	//
	// This configuration provides customization to how the model identifies and extracts
	// relevant information for memory storage. You can use the default user prompt or create a customized one.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/system-prompts.html
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AppendToPrompt *string `field:"required" json:"appendToPrompt" yaml:"appendToPrompt"`
	// The model to use for consolidation/extraction.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Model awsbedrockalpha.IBedrockInvokable `field:"required" json:"model" yaml:"model"`
}

