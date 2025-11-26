package previewawsbedrockagentcoremixins


// The memory override configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userPreferenceOverrideExtractionConfigurationInputProperty := &UserPreferenceOverrideExtractionConfigurationInputProperty{
//   	AppendToPrompt: jsii.String("appendToPrompt"),
//   	ModelId: jsii.String("modelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferenceoverrideextractionconfigurationinput.html
//
type CfnMemoryPropsMixin_UserPreferenceOverrideExtractionConfigurationInputProperty struct {
	// The extraction configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferenceoverrideextractionconfigurationinput.html#cfn-bedrockagentcore-memory-userpreferenceoverrideextractionconfigurationinput-appendtoprompt
	//
	AppendToPrompt *string `field:"optional" json:"appendToPrompt" yaml:"appendToPrompt"`
	// The memory override for the model ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferenceoverrideextractionconfigurationinput.html#cfn-bedrockagentcore-memory-userpreferenceoverrideextractionconfigurationinput-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}

