package awsbedrock


// Contains details about a variant of the prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptVariantProperty := &PromptVariantProperty{
//   	Name: jsii.String("name"),
//   	TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   		Text: &TextPromptTemplateConfigurationProperty{
//   			InputVariables: []interface{}{
//   				&PromptInputVariableProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Text: jsii.String("text"),
//   			TextS3Location: &TextS3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				Version: jsii.String("version"),
//   			},
//   		},
//   	},
//   	TemplateType: jsii.String("templateType"),
//
//   	// the properties below are optional
//   	InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   		Text: &PromptModelInferenceConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   			StopSequences: []*string{
//   				jsii.String("stopSequences"),
//   			},
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   	},
//   	ModelId: jsii.String("modelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html
//
type CfnPrompt_PromptVariantProperty struct {
	// The name of the prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains configurations for the prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"required" json:"templateConfiguration" yaml:"templateConfiguration"`
	// The type of prompt template to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-templatetype
	//
	TemplateType *string `field:"required" json:"templateType" yaml:"templateType"`
	// Contains inference configurations for the prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-inferenceconfiguration
	//
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// The unique identifier of the model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) with which to run inference on the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}

