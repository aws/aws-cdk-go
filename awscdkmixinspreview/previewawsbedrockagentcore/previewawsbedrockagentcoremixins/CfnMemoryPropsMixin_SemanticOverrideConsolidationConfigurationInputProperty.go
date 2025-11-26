package previewawsbedrockagentcoremixins


// The memory override configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   semanticOverrideConsolidationConfigurationInputProperty := &SemanticOverrideConsolidationConfigurationInputProperty{
//   	AppendToPrompt: jsii.String("appendToPrompt"),
//   	ModelId: jsii.String("modelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticoverrideconsolidationconfigurationinput.html
//
type CfnMemoryPropsMixin_SemanticOverrideConsolidationConfigurationInputProperty struct {
	// The override configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticoverrideconsolidationconfigurationinput.html#cfn-bedrockagentcore-memory-semanticoverrideconsolidationconfigurationinput-appendtoprompt
	//
	AppendToPrompt *string `field:"optional" json:"appendToPrompt" yaml:"appendToPrompt"`
	// The memory override model ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticoverrideconsolidationconfigurationinput.html#cfn-bedrockagentcore-memory-semanticoverrideconsolidationconfigurationinput-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}

