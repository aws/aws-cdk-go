package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a guardrail to detect and filter harmful content in your generative AI application.
//
// Amazon Bedrock Guardrails provides the following safeguards (also known as policies) to detect and filter harmful content:
//
// - *Content filters* - Detect and filter harmful text or image content in input prompts or model responses. Filtering is done based on detection of certain predefined harmful content categories: Hate, Insults, Sexual, Violence, Misconduct and Prompt Attack. You also can adjust the filter strength for each of these categories.
// - *Denied topics* - Define a set of topics that are undesirable in the context of your application. The filter will help block them if detected in user queries or model responses.
// - *Word filters* - Configure filters to help block undesirable words, phrases, and profanity (exact match). Such words can include offensive terms, competitor names, etc.
// - *Sensitive information filters* - Configure filters to help block or mask sensitive information, such as personally identifiable information (PII), or custom regex in user inputs and model responses. Blocking or masking is done based on probabilistic detection of sensitive information in standard formats in entities such as SSN number, Date of Birth, address, etc. This also allows configuring regular expression based detection of patterns for identifiers.
// - *Contextual grounding check* - Help detect and filter hallucinations in model responses based on grounding in a source and relevance to the user query.
//
// For more information, see [How Amazon Bedrock Guardrails works](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-how.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGuardrailPropsMixin := awscdkmixinspreview.Mixins.NewCfnGuardrailPropsMixin(&CfnGuardrailMixinProps{
//   	AutomatedReasoningPolicyConfig: &AutomatedReasoningPolicyConfigProperty{
//   		ConfidenceThreshold: jsii.Number(123),
//   		Policies: []*string{
//   			jsii.String("policies"),
//   		},
//   	},
//   	BlockedInputMessaging: jsii.String("blockedInputMessaging"),
//   	BlockedOutputsMessaging: jsii.String("blockedOutputsMessaging"),
//   	ContentPolicyConfig: &ContentPolicyConfigProperty{
//   		ContentFiltersTierConfig: &ContentFiltersTierConfigProperty{
//   			TierName: jsii.String("tierName"),
//   		},
//   		FiltersConfig: []interface{}{
//   			&ContentFilterConfigProperty{
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				InputModalities: []*string{
//   					jsii.String("inputModalities"),
//   				},
//   				InputStrength: jsii.String("inputStrength"),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				OutputModalities: []*string{
//   					jsii.String("outputModalities"),
//   				},
//   				OutputStrength: jsii.String("outputStrength"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ContextualGroundingPolicyConfig: &ContextualGroundingPolicyConfigProperty{
//   		FiltersConfig: []interface{}{
//   			&ContextualGroundingFilterConfigProperty{
//   				Action: jsii.String("action"),
//   				Enabled: jsii.Boolean(false),
//   				Threshold: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	CrossRegionConfig: &GuardrailCrossRegionConfigProperty{
//   		GuardrailProfileArn: jsii.String("guardrailProfileArn"),
//   	},
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Name: jsii.String("name"),
//   	SensitiveInformationPolicyConfig: &SensitiveInformationPolicyConfigProperty{
//   		PiiEntitiesConfig: []interface{}{
//   			&PiiEntityConfigProperty{
//   				Action: jsii.String("action"),
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		RegexesConfig: []interface{}{
//   			&RegexConfigProperty{
//   				Action: jsii.String("action"),
//   				Description: jsii.String("description"),
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TopicPolicyConfig: &TopicPolicyConfigProperty{
//   		TopicsConfig: []interface{}{
//   			&TopicConfigProperty{
//   				Definition: jsii.String("definition"),
//   				Examples: []*string{
//   					jsii.String("examples"),
//   				},
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		TopicsTierConfig: &TopicsTierConfigProperty{
//   			TierName: jsii.String("tierName"),
//   		},
//   	},
//   	WordPolicyConfig: &WordPolicyConfigProperty{
//   		ManagedWordListsConfig: []interface{}{
//   			&ManagedWordsConfigProperty{
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		WordsConfig: []interface{}{
//   			&WordConfigProperty{
//   				InputAction: jsii.String("inputAction"),
//   				InputEnabled: jsii.Boolean(false),
//   				OutputAction: jsii.String("outputAction"),
//   				OutputEnabled: jsii.Boolean(false),
//   				Text: jsii.String("text"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html
//
type CfnGuardrailPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGuardrailMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGuardrailPropsMixin
type jsiiProxy_CfnGuardrailPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGuardrailPropsMixin) Props() *CfnGuardrailMixinProps {
	var returns *CfnGuardrailMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrailPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::Guardrail`.
func NewCfnGuardrailPropsMixin(props *CfnGuardrailMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGuardrailPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGuardrailPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGuardrailPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnGuardrailPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::Guardrail`.
func NewCfnGuardrailPropsMixin_Override(c CfnGuardrailPropsMixin, props *CfnGuardrailMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnGuardrailPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGuardrailPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGuardrailPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnGuardrailPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGuardrailPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnGuardrailPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGuardrailPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGuardrailPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

