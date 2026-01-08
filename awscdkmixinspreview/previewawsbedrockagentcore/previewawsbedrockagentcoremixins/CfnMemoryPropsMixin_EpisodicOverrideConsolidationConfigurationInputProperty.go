package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   episodicOverrideConsolidationConfigurationInputProperty := &EpisodicOverrideConsolidationConfigurationInputProperty{
//   	AppendToPrompt: jsii.String("appendToPrompt"),
//   	ModelId: jsii.String("modelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverrideconsolidationconfigurationinput.html
//
type CfnMemoryPropsMixin_EpisodicOverrideConsolidationConfigurationInputProperty struct {
	// Text prompt for model instructions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverrideconsolidationconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverrideconsolidationconfigurationinput-appendtoprompt
	//
	AppendToPrompt *string `field:"optional" json:"appendToPrompt" yaml:"appendToPrompt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverrideconsolidationconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverrideconsolidationconfigurationinput-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}

