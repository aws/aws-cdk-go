package awsbedrockalpha


// GuardrailCrossRegionConfigProperty.
//
// Example:
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   	Description: jsii.String("Guardrail with cross-region configuration for enhanced language support"),
//   	CrossRegionConfig: &GuardrailCrossRegionConfigProperty{
//   		GuardrailProfileArn: jsii.String("arn:aws:bedrock:us-east-1:123456789012:guardrail-profile/my-profile"),
//   	},
//   	// Use STANDARD tier for enhanced capabilities
//   	ContentFiltersTierConfig: bedrock.TierConfig_STANDARD,
//   	TopicsTierConfig: bedrock.TierConfig_STANDARD,
//   })
//
// Experimental.
type GuardrailCrossRegionConfigProperty struct {
	// The arn of thesystem-defined guardrail profile that you're using with your guardrail.
	//
	// Guardrail profiles define the destination AWS Regions where guardrail inference requests can be automatically routed.
	// Using guardrail profiles helps maintain guardrail performance and reliability when demand increases.
	// Default: - No cross-region configuration.
	//
	// Experimental.
	GuardrailProfileArn *string `field:"required" json:"guardrailProfileArn" yaml:"guardrailProfileArn"`
}

