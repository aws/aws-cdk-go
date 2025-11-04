package awsbedrockalpha


// Properties for creating a Cross-Region Inference Profile.
//
// Example:
//   // Create a cross-region inference profile
//   crossRegionProfile := bedrock.CrossRegionInferenceProfile_FromConfig(&CrossRegionInferenceProfileProps{
//   	GeoRegion: bedrock.CrossRegionInferenceProfileRegion_US,
//   	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
//   })
//
//   // Use the cross-region profile with an agent
//   agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
//   	FoundationModel: crossRegionProfile,
//   	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about agriculture."),
//   })
//
// Experimental.
type CrossRegionInferenceProfileProps struct {
	// The geographic region where the traffic is going to be distributed.
	//
	// Routing
	// factors in user traffic, demand and utilization of resources.
	// Experimental.
	GeoRegion CrossRegionInferenceProfileRegion `field:"required" json:"geoRegion" yaml:"geoRegion"`
	// A foundation model supporting cross-region inference.
	//
	// The model must have cross-region support enabled.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference-support.html
	//
	// Experimental.
	Model BedrockFoundationModel `field:"required" json:"model" yaml:"model"`
}

