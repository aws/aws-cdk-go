package awsbedrockalpha


// Properties for creating an Application Inference Profile.
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
type ApplicationInferenceProfileProps struct {
	// The name of the application inference profile.
	//
	// This name will be used to identify the inference profile in the AWS console and APIs.
	// - Required:  Yes
	// - Maximum length: 64 characters
	// - Pattern: `^([0-9a-zA-Z:.][ _-]?)+$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-inferenceprofilename
	//
	// Experimental.
	ApplicationInferenceProfileName *string `field:"required" json:"applicationInferenceProfileName" yaml:"applicationInferenceProfileName"`
	// The model source for this inference profile.
	//
	// To create an application inference profile for one Region, specify a foundation model.
	// Usage and costs for requests made to that Region with that model will be tracked.
	//
	// To create an application inference profile for multiple Regions,
	// specify a cross region (system-defined) inference profile.
	// The inference profile will route requests to the Regions defined in
	// the cross region (system-defined) inference profile that you choose.
	// Usage and costs for requests made to the Regions in the inference profile will be tracked.
	// Experimental.
	ModelSource IBedrockInvokable `field:"required" json:"modelSource" yaml:"modelSource"`
	// Description of the inference profile. Provides additional context about the purpose and usage of this inference profile.
	//
	// - Maximum length: 200 characters when provided
	// - Pattern: `^([0-9a-zA-Z:.][ _-]?)+$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-description
	//
	// Default: - No description is provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of tags associated with the inference profile.
	//
	// Tags help you organize and categorize your AWS resources.
	// Default: - No tags are applied.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

