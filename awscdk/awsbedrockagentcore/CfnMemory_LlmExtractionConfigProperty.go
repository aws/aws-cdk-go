package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   llmExtractionConfigProperty := &LlmExtractionConfigProperty{
//   	Definition: jsii.String("definition"),
//
//   	// the properties below are optional
//   	LlmExtractionInstruction: jsii.String("llmExtractionInstruction"),
//   	Validation: &ValidationProperty{
//   		NumberValidation: &NumberValidationProperty{
//   			MaxValue: jsii.Number(123),
//   			MinValue: jsii.Number(123),
//   		},
//   		StringListValidation: &StringListValidationProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			MaxItems: jsii.Number(123),
//   		},
//   		StringValidation: &StringValidationProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-llmextractionconfig.html
//
type CfnMemory_LlmExtractionConfigProperty struct {
	// Definition for the metadata schema entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-llmextractionconfig.html#cfn-bedrockagentcore-memory-llmextractionconfig-definition
	//
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// LLM extraction instruction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-llmextractionconfig.html#cfn-bedrockagentcore-memory-llmextractionconfig-llmextractioninstruction
	//
	LlmExtractionInstruction *string `field:"optional" json:"llmExtractionInstruction" yaml:"llmExtractionInstruction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-llmextractionconfig.html#cfn-bedrockagentcore-memory-llmextractionconfig-validation
	//
	Validation interface{} `field:"optional" json:"validation" yaml:"validation"`
}

