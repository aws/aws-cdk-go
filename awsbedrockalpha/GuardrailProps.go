package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Guardrail.
//
// Example:
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   })
//   // Add regex filter with input/output actions
//   guardrail.AddRegexFilter(&RegexFilter{
//   	Name: jsii.String("TestRegexFilter"),
//   	Pattern: jsii.String("test-pattern"),
//   	Action: bedrock.GuardrailAction_ANONYMIZE,
//   	// below props are optional
//   	Description: jsii.String("This is a test regex filter"),
//   	InputAction: bedrock.GuardrailAction_BLOCK,
//   	InputEnabled: jsii.Boolean(true),
//   	OutputAction: bedrock.GuardrailAction_ANONYMIZE,
//   	OutputEnabled: jsii.Boolean(true),
//   })
//
// Experimental.
type GuardrailProps struct {
	// The name of the guardrail.
	//
	// This will be used as the physical name of the guardrail.
	// Experimental.
	GuardrailName *string `field:"required" json:"guardrailName" yaml:"guardrailName"`
	// The message to return when the guardrail blocks a prompt.
	//
	// Must be between 1 and 500 characters.
	// Default: "Sorry, your query violates our usage policy."
	//
	// Experimental.
	BlockedInputMessaging *string `field:"optional" json:"blockedInputMessaging" yaml:"blockedInputMessaging"`
	// The message to return when the guardrail blocks a model response.
	//
	// Must be between 1 and 500 characters.
	// Default: "Sorry, I am unable to answer your question because of our usage policy."
	//
	// Experimental.
	BlockedOutputsMessaging *string `field:"optional" json:"blockedOutputsMessaging" yaml:"blockedOutputsMessaging"`
	// The content filters to apply to the guardrail.
	// Default: [].
	//
	// Experimental.
	ContentFilters *[]*ContentFilter `field:"optional" json:"contentFilters" yaml:"contentFilters"`
	// The tier configuration to apply to the guardrail.
	// Default: filters.TierConfig.CLASSIC
	//
	// Experimental.
	ContentFiltersTierConfig TierConfig `field:"optional" json:"contentFiltersTierConfig" yaml:"contentFiltersTierConfig"`
	// The contextual grounding filters to apply to the guardrail.
	// Default: [].
	//
	// Experimental.
	ContextualGroundingFilters *[]*ContextualGroundingFilter `field:"optional" json:"contextualGroundingFilters" yaml:"contextualGroundingFilters"`
	// The cross-region configuration for the guardrail.
	//
	// This is optional and when provided, it should be of type GuardrailCrossRegionConfigProperty.
	// Default: - No cross-region configuration.
	//
	// Experimental.
	CrossRegionConfig *GuardrailCrossRegionConfigProperty `field:"optional" json:"crossRegionConfig" yaml:"crossRegionConfig"`
	// A list of policies related to topics that the guardrail should deny.
	// Default: [].
	//
	// Experimental.
	DeniedTopics *[]Topic `field:"optional" json:"deniedTopics" yaml:"deniedTopics"`
	// The description of the guardrail.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A custom KMS key to use for encrypting data.
	// Default: - Data is encrypted by default with a key that AWS owns and manages for you.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The managed word filters to apply to the guardrail.
	// Default: [].
	//
	// Experimental.
	ManagedWordListFilters *[]*ManagedWordFilter `field:"optional" json:"managedWordListFilters" yaml:"managedWordListFilters"`
	// The PII filters to apply to the guardrail.
	// Default: [].
	//
	// Experimental.
	PiiFilters *[]*PIIFilter `field:"optional" json:"piiFilters" yaml:"piiFilters"`
	// The regular expression (regex) filters to apply to the guardrail.
	// Default: [].
	//
	// Experimental.
	RegexFilters *[]*RegexFilter `field:"optional" json:"regexFilters" yaml:"regexFilters"`
	// The tier configuration to apply to the guardrail.
	// Default: filters.TierConfig.CLASSIC
	//
	// Experimental.
	TopicsTierConfig TierConfig `field:"optional" json:"topicsTierConfig" yaml:"topicsTierConfig"`
	// The word filters to apply to the guardrail.
	// Default: [].
	//
	// Experimental.
	WordFilters *[]*WordFilter `field:"optional" json:"wordFilters" yaml:"wordFilters"`
}

