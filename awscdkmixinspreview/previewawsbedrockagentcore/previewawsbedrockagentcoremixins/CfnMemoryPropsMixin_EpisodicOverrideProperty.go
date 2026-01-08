package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   episodicOverrideProperty := &EpisodicOverrideProperty{
//   	Consolidation: &EpisodicOverrideConsolidationConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   	Extraction: &EpisodicOverrideExtractionConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   	Reflection: &EpisodicOverrideReflectionConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverride.html
//
type CfnMemoryPropsMixin_EpisodicOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverride.html#cfn-bedrockagentcore-memory-episodicoverride-consolidation
	//
	Consolidation interface{} `field:"optional" json:"consolidation" yaml:"consolidation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverride.html#cfn-bedrockagentcore-memory-episodicoverride-extraction
	//
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverride.html#cfn-bedrockagentcore-memory-episodicoverride-reflection
	//
	Reflection interface{} `field:"optional" json:"reflection" yaml:"reflection"`
}

