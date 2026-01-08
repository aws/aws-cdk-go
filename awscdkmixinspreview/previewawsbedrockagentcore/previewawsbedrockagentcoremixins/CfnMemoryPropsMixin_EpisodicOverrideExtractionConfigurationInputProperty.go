package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   episodicOverrideExtractionConfigurationInputProperty := &EpisodicOverrideExtractionConfigurationInputProperty{
//   	AppendToPrompt: jsii.String("appendToPrompt"),
//   	ModelId: jsii.String("modelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverrideextractionconfigurationinput.html
//
type CfnMemoryPropsMixin_EpisodicOverrideExtractionConfigurationInputProperty struct {
	// Text prompt for model instructions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverrideextractionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverrideextractionconfigurationinput-appendtoprompt
	//
	AppendToPrompt *string `field:"optional" json:"appendToPrompt" yaml:"appendToPrompt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverrideextractionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverrideextractionconfigurationinput-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}

