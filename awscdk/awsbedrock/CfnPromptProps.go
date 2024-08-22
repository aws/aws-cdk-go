package awsbedrock


// Properties for defining a `CfnPrompt`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPromptProps := &CfnPromptProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CustomerEncryptionKeyArn: jsii.String("customerEncryptionKeyArn"),
//   	DefaultVariant: jsii.String("defaultVariant"),
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Variants: []interface{}{
//   		&PromptVariantProperty{
//   			Name: jsii.String("name"),
//   			TemplateType: jsii.String("templateType"),
//
//   			// the properties below are optional
//   			InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   				Text: &PromptModelInferenceConfigurationProperty{
//   					MaxTokens: jsii.Number(123),
//   					StopSequences: []*string{
//   						jsii.String("stopSequences"),
//   					},
//   					Temperature: jsii.Number(123),
//   					TopK: jsii.Number(123),
//   					TopP: jsii.Number(123),
//   				},
//   			},
//   			ModelId: jsii.String("modelId"),
//   			TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   				Text: &TextPromptTemplateConfigurationProperty{
//   					InputVariables: []interface{}{
//   						&PromptInputVariableProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Text: jsii.String("text"),
//   					TextS3Location: &TextS3LocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//
//   						// the properties below are optional
//   						Version: jsii.String("version"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-prompt.html
//
type CfnPromptProps struct {
	// The name of the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-prompt.html#cfn-bedrock-prompt-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the KMS key that the prompt is encrypted with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-prompt.html#cfn-bedrock-prompt-customerencryptionkeyarn
	//
	CustomerEncryptionKeyArn *string `field:"optional" json:"customerEncryptionKeyArn" yaml:"customerEncryptionKeyArn"`
	// The name of the default variant for the prompt.
	//
	// This value must match the `name` field in the relevant [PromptVariant](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_PromptVariant.html) object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-prompt.html#cfn-bedrock-prompt-defaultvariant
	//
	DefaultVariant *string `field:"optional" json:"defaultVariant" yaml:"defaultVariant"`
	// The description of the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-prompt.html#cfn-bedrock-prompt-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-prompt.html#cfn-bedrock-prompt-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// A list of objects, each containing details about a variant of the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-prompt.html#cfn-bedrock-prompt-variants
	//
	Variants interface{} `field:"optional" json:"variants" yaml:"variants"`
}

