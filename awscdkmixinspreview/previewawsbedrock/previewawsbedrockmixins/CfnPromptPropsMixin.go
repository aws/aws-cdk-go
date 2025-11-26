package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a prompt in your prompt library that you can add to a flow.
//
// For more information, see [Prompt management in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management.html) , [Create a prompt using Prompt management](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-create.html) and [Prompt flows in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/flows.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//   var any interface{}
//   var auto interface{}
//   var json interface{}
//
//   cfnPromptPropsMixin := awscdkmixinspreview.Mixins.NewCfnPromptPropsMixin(&CfnPromptMixinProps{
//   	CustomerEncryptionKeyArn: jsii.String("customerEncryptionKeyArn"),
//   	DefaultVariant: jsii.String("defaultVariant"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Variants: []interface{}{
//   		&PromptVariantProperty{
//   			AdditionalModelRequestFields: additionalModelRequestFields,
//   			GenAiResource: &PromptGenAiResourceProperty{
//   				Agent: &PromptAgentResourceProperty{
//   					AgentIdentifier: jsii.String("agentIdentifier"),
//   				},
//   			},
//   			InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   				Text: &PromptModelInferenceConfigurationProperty{
//   					MaxTokens: jsii.Number(123),
//   					StopSequences: []*string{
//   						jsii.String("stopSequences"),
//   					},
//   					Temperature: jsii.Number(123),
//   					TopP: jsii.Number(123),
//   				},
//   			},
//   			Metadata: []interface{}{
//   				&PromptMetadataEntryProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			ModelId: jsii.String("modelId"),
//   			Name: jsii.String("name"),
//   			TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   				Chat: &ChatPromptTemplateConfigurationProperty{
//   					InputVariables: []interface{}{
//   						&PromptInputVariableProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Messages: []interface{}{
//   						&MessageProperty{
//   							Content: []interface{}{
//   								&ContentBlockProperty{
//   									CachePoint: &CachePointBlockProperty{
//   										Type: jsii.String("type"),
//   									},
//   									Text: jsii.String("text"),
//   								},
//   							},
//   							Role: jsii.String("role"),
//   						},
//   					},
//   					System: []interface{}{
//   						&SystemContentBlockProperty{
//   							CachePoint: &CachePointBlockProperty{
//   								Type: jsii.String("type"),
//   							},
//   							Text: jsii.String("text"),
//   						},
//   					},
//   					ToolConfiguration: &ToolConfigurationProperty{
//   						ToolChoice: &ToolChoiceProperty{
//   							Any: any,
//   							Auto: auto,
//   							Tool: &SpecificToolChoiceProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Tools: []interface{}{
//   							&ToolProperty{
//   								CachePoint: &CachePointBlockProperty{
//   									Type: jsii.String("type"),
//   								},
//   								ToolSpec: &ToolSpecificationProperty{
//   									Description: jsii.String("description"),
//   									InputSchema: &ToolInputSchemaProperty{
//   										Json: json,
//   									},
//   									Name: jsii.String("name"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				Text: &TextPromptTemplateConfigurationProperty{
//   					CachePoint: &CachePointBlockProperty{
//   						Type: jsii.String("type"),
//   					},
//   					InputVariables: []interface{}{
//   						&PromptInputVariableProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Text: jsii.String("text"),
//   					TextS3Location: &TextS3LocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//   						Version: jsii.String("version"),
//   					},
//   				},
//   			},
//   			TemplateType: jsii.String("templateType"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-prompt.html
//
type CfnPromptPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPromptMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPromptPropsMixin
type jsiiProxy_CfnPromptPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPromptPropsMixin) Props() *CfnPromptMixinProps {
	var returns *CfnPromptMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPromptPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::Prompt`.
func NewCfnPromptPropsMixin(props *CfnPromptMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPromptPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPromptPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPromptPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnPromptPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::Prompt`.
func NewCfnPromptPropsMixin_Override(c CfnPromptPropsMixin, props *CfnPromptMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnPromptPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPromptPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPromptPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnPromptPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPromptPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnPromptPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPromptPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPromptPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

