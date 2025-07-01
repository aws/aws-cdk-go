package awswisdom


// Properties for defining a `CfnAIPrompt`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAIPromptProps := &CfnAIPromptProps{
//   	ApiFormat: jsii.String("apiFormat"),
//   	ModelId: jsii.String("modelId"),
//   	TemplateConfiguration: &AIPromptTemplateConfigurationProperty{
//   		TextFullAiPromptEditTemplateConfiguration: &TextFullAIPromptEditTemplateConfigurationProperty{
//   			Text: jsii.String("text"),
//   		},
//   	},
//   	TemplateType: jsii.String("templateType"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AssistantId: jsii.String("assistantId"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html
//
type CfnAIPromptProps struct {
	// The API format used for this AI Prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html#cfn-wisdom-aiprompt-apiformat
	//
	ApiFormat *string `field:"required" json:"apiFormat" yaml:"apiFormat"`
	// The identifier of the model used for this AI Prompt. The following model Ids are supported:.
	//
	// - `anthropic.claude-3-haiku--v1:0`
	// - `apac.amazon.nova-lite-v1:0`
	// - `apac.amazon.nova-micro-v1:0`
	// - `apac.amazon.nova-pro-v1:0`
	// - `apac.anthropic.claude-3-5-sonnet--v2:0`
	// - `apac.anthropic.claude-3-haiku-20240307-v1:0`
	// - `eu.amazon.nova-lite-v1:0`
	// - `eu.amazon.nova-micro-v1:0`
	// - `eu.amazon.nova-pro-v1:0`
	// - `eu.anthropic.claude-3-7-sonnet-20250219-v1:0`
	// - `eu.anthropic.claude-3-haiku-20240307-v1:0`
	// - `us.amazon.nova-lite-v1:0`
	// - `us.amazon.nova-micro-v1:0`
	// - `us.amazon.nova-pro-v1:0`
	// - `us.anthropic.claude-3-5-haiku-20241022-v1:0`
	// - `us.anthropic.claude-3-7-sonnet-20250219-v1:0`
	// - `us.anthropic.claude-3-haiku-20240307-v1:0`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html#cfn-wisdom-aiprompt-modelid
	//
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// The configuration of the prompt template for this AI Prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html#cfn-wisdom-aiprompt-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"required" json:"templateConfiguration" yaml:"templateConfiguration"`
	// The type of the prompt template for this AI Prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html#cfn-wisdom-aiprompt-templatetype
	//
	TemplateType *string `field:"required" json:"templateType" yaml:"templateType"`
	// The type of this AI Prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html#cfn-wisdom-aiprompt-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The identifier of the Amazon Q in Connect assistant.
	//
	// Can be either the ID or the ARN. URLs cannot contain the ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html#cfn-wisdom-aiprompt-assistantid
	//
	AssistantId *string `field:"optional" json:"assistantId" yaml:"assistantId"`
	// The description of the AI Prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html#cfn-wisdom-aiprompt-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the AI Prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html#cfn-wisdom-aiprompt-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiprompt.html#cfn-wisdom-aiprompt-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

