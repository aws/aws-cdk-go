package awsbedrockalpha


// Geographic regions supported for cross-region inference profiles.
//
// These regions help distribute traffic across multiple AWS regions for better
// throughput and resilience during peak demands.
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
type CrossRegionInferenceProfileRegion string

const (
	// Cross-region Inference Identifier for the European area.
	//
	// According to the model chosen, this might include:
	// - Frankfurt (`eu-central-1`)
	// - Ireland (`eu-west-1`)
	// - Paris (`eu-west-3`).
	// Experimental.
	CrossRegionInferenceProfileRegion_EU CrossRegionInferenceProfileRegion = "EU"
	// Cross-region Inference Identifier for the United States area.
	//
	// According to the model chosen, this might include:
	// - N. Virginia (`us-east-1`)
	// - Oregon (`us-west-2`)
	// - Ohio (`us-east-2`).
	// Experimental.
	CrossRegionInferenceProfileRegion_US CrossRegionInferenceProfileRegion = "US"
	// Cross-region Inference Identifier for the Asia-Pacific area.
	//
	// According to the model chosen, this might include:
	// - Tokyo (`ap-northeast-1`)
	// - Seoul (`ap-northeast-2`)
	// - Mumbai (`ap-south-1`)
	// - Singapore (`ap-southeast-1`)
	// - Sydney (`ap-southeast-2`).
	// Experimental.
	CrossRegionInferenceProfileRegion_APAC CrossRegionInferenceProfileRegion = "APAC"
)

