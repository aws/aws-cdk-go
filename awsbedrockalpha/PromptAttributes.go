package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Attributes for specifying an imported Bedrock Prompt.
//
// Example:
//   // Import an existing prompt by ARN
//   importedPrompt := bedrock.Prompt_FromPromptAttributes(this, jsii.String("ImportedPrompt"), &PromptAttributes{
//   	PromptArn: jsii.String("arn:aws:bedrock:region:account:prompt/prompt-id"),
//   	KmsKey: kms.Key_FromKeyArn(this, jsii.String("ImportedKey"), jsii.String("arn:aws:kms:region:account:key/key-id")),
//   	 // optional
//   	PromptVersion: jsii.String("1"),
//   })
//
// Experimental.
type PromptAttributes struct {
	// The ARN of the prompt.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:prompt/PROMPT12345"
	//
	// Experimental.
	PromptArn *string `field:"required" json:"promptArn" yaml:"promptArn"`
	// Optional KMS encryption key associated with this prompt.
	// Default: undefined - An AWS managed key is used.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The version of the prompt.
	// Default: "DRAFT".
	//
	// Experimental.
	PromptVersion *string `field:"optional" json:"promptVersion" yaml:"promptVersion"`
}

