package previewawsbedrockagentcoremixins


// The memory override.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   semanticOverrideProperty := &SemanticOverrideProperty{
//   	Consolidation: &SemanticOverrideConsolidationConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   	Extraction: &SemanticOverrideExtractionConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticoverride.html
//
type CfnMemoryPropsMixin_SemanticOverrideProperty struct {
	// The memory override consolidation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticoverride.html#cfn-bedrockagentcore-memory-semanticoverride-consolidation
	//
	Consolidation interface{} `field:"optional" json:"consolidation" yaml:"consolidation"`
	// The memory override extraction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticoverride.html#cfn-bedrockagentcore-memory-semanticoverride-extraction
	//
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
}

