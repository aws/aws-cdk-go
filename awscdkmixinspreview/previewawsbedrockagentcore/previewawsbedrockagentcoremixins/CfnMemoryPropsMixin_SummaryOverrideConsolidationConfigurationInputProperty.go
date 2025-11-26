package previewawsbedrockagentcoremixins


// The consolidation configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   summaryOverrideConsolidationConfigurationInputProperty := &SummaryOverrideConsolidationConfigurationInputProperty{
//   	AppendToPrompt: jsii.String("appendToPrompt"),
//   	ModelId: jsii.String("modelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summaryoverrideconsolidationconfigurationinput.html
//
type CfnMemoryPropsMixin_SummaryOverrideConsolidationConfigurationInputProperty struct {
	// The memory override configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summaryoverrideconsolidationconfigurationinput.html#cfn-bedrockagentcore-memory-summaryoverrideconsolidationconfigurationinput-appendtoprompt
	//
	AppendToPrompt *string `field:"optional" json:"appendToPrompt" yaml:"appendToPrompt"`
	// The memory override configuration model ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summaryoverrideconsolidationconfigurationinput.html#cfn-bedrockagentcore-memory-summaryoverrideconsolidationconfigurationinput-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}

