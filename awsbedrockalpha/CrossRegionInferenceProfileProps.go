package awsbedrockalpha


// Properties for creating a Cross-Region Inference Profile.
//
// Example:
//   // Create a cross-region inference profile
//   crossRegionProfile := bedrock.CrossRegionInferenceProfile_FromConfig(&CrossRegionInferenceProfileProps{
//   	GeoRegion: bedrock.CrossRegionInferenceProfileRegion_US,
//   	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V2_0(),
//   })
//
//   // Create an application inference profile across regions
//   appProfile := bedrock.NewApplicationInferenceProfile(this, jsii.String("MyMultiRegionProfile"), &ApplicationInferenceProfileProps{
//   	ApplicationInferenceProfileName: jsii.String("claude-35-sonnet-v2-multi-region"),
//   	ModelSource: crossRegionProfile,
//   	Description: jsii.String("Multi-region application profile for cost tracking"),
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

