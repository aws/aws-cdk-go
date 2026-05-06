package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   extractionConfigProperty := &ExtractionConfigProperty{
//   	LlmExtractionConfig: &LlmExtractionConfigProperty{
//   		Definition: jsii.String("definition"),
//   		LlmExtractionInstruction: jsii.String("llmExtractionInstruction"),
//   		Validation: &ValidationProperty{
//   			NumberValidation: &NumberValidationProperty{
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//   			},
//   			StringListValidation: &StringListValidationProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				MaxItems: jsii.Number(123),
//   			},
//   			StringValidation: &StringValidationProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-extractionconfig.html
//
type CfnMemoryPropsMixin_ExtractionConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-extractionconfig.html#cfn-bedrockagentcore-memory-extractionconfig-llmextractionconfig
	//
	LlmExtractionConfig interface{} `field:"optional" json:"llmExtractionConfig" yaml:"llmExtractionConfig"`
}

