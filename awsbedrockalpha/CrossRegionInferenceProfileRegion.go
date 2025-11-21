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
	// Global cross-region Inference Identifier.
	//
	// Routes requests to any supported commercial AWS Region.
	// Experimental.
	CrossRegionInferenceProfileRegion_GLOBAL CrossRegionInferenceProfileRegion = "GLOBAL"
	// Cross-region Inference Identifier for the European area.
	//
	// According to the model chosen, this might include:
	// - Frankfurt (`eu-central-1`)
	// - Ireland (`eu-west-1`)
	// - Paris (`eu-west-3`)
	// - London (`eu-west-2`)
	// - Stockholm (`eu-north-1`)
	// - Milan (`eu-south-1`)
	// - Spain (`eu-south-2`)
	// - Zurich (`eu-central-2`).
	// Experimental.
	CrossRegionInferenceProfileRegion_EU CrossRegionInferenceProfileRegion = "EU"
	// Cross-region Inference Identifier for the United States area.
	//
	// According to the model chosen, this might include:
	// - N. Virginia (`us-east-1`)
	// - Ohio (`us-east-2`)
	// - Oregon (`us-west-2`).
	// Experimental.
	CrossRegionInferenceProfileRegion_US CrossRegionInferenceProfileRegion = "US"
	// Cross-region Inference Identifier for the US GovCloud area.
	//
	// According to the model chosen, this might include:
	// - GovCloud US-East (`us-gov-east-1`)
	// - GovCloud US-West (`us-gov-west-1`).
	// Experimental.
	CrossRegionInferenceProfileRegion_US_GOV CrossRegionInferenceProfileRegion = "US_GOV"
	// Cross-region Inference Identifier for the Asia-Pacific area.
	//
	// According to the model chosen, this might include:
	// - Tokyo (`ap-northeast-1`)
	// - Seoul (`ap-northeast-2`)
	// - Osaka (`ap-northeast-3`)
	// - Mumbai (`ap-south-1`)
	// - Hyderabad (`ap-south-2`)
	// - Singapore (`ap-southeast-1`)
	// - Sydney (`ap-southeast-2`)
	// - Jakarta (`ap-southeast-3`)
	// - Melbourne (`ap-southeast-4`)
	// - Malaysia (`ap-southeast-5`)
	// - Thailand (`ap-southeast-7`)
	// - Taipei (`ap-east-2`)
	// - Middle East (UAE) (`me-central-1`).
	// Experimental.
	CrossRegionInferenceProfileRegion_APAC CrossRegionInferenceProfileRegion = "APAC"
	// Cross-region Inference Identifier for the Japan area.
	//
	// According to the model chosen, this might include:
	// - Tokyo (`ap-northeast-1`)
	// - Osaka (`ap-northeast-3`).
	// Experimental.
	CrossRegionInferenceProfileRegion_JP CrossRegionInferenceProfileRegion = "JP"
	// Cross-region Inference Identifier for the Australia area.
	//
	// According to the model chosen, this might include:
	// - Sydney (`ap-southeast-2`)
	// - Melbourne (`ap-southeast-4`).
	// Experimental.
	CrossRegionInferenceProfileRegion_AU CrossRegionInferenceProfileRegion = "AU"
)

