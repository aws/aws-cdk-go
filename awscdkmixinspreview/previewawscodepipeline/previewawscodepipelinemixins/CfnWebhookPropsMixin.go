package previewawscodepipelinemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodepipeline/previewawscodepipelinemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodePipeline::Webhook` resource creates and registers your webhook.
//
// After the webhook is created and registered, it triggers your pipeline to start every time an external event occurs. For more information, see [Migrate polling pipelines to use event-based change detection](https://docs.aws.amazon.com/codepipeline/latest/userguide/update-change-detection.html) in the *AWS CodePipeline User Guide* .
//
// We strongly recommend that you use AWS Secrets Manager to store your credentials. If you use Secrets Manager, you must have already configured and stored your secret parameters in Secrets Manager. For more information, see [Using Dynamic References to Specify Template Values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) .
//
// > When passing secret parameters, do not enter the value directly into the template. The value is rendered as plaintext and is therefore readable. For security reasons, do not use plaintext in your AWS CloudFormation template to store your credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebhookPropsMixin := awscdkmixinspreview.Mixins.NewCfnWebhookPropsMixin(&CfnWebhookMixinProps{
//   	Authentication: jsii.String("authentication"),
//   	AuthenticationConfiguration: &WebhookAuthConfigurationProperty{
//   		AllowedIpRange: jsii.String("allowedIpRange"),
//   		SecretToken: jsii.String("secretToken"),
//   	},
//   	Filters: []interface{}{
//   		&WebhookFilterRuleProperty{
//   			JsonPath: jsii.String("jsonPath"),
//   			MatchEquals: jsii.String("matchEquals"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RegisterWithThirdParty: jsii.Boolean(false),
//   	TargetAction: jsii.String("targetAction"),
//   	TargetPipeline: jsii.String("targetPipeline"),
//   	TargetPipelineVersion: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html
//
type CfnWebhookPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWebhookMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWebhookPropsMixin
type jsiiProxy_CfnWebhookPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWebhookPropsMixin) Props() *CfnWebhookMixinProps {
	var returns *CfnWebhookMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhookPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodePipeline::Webhook`.
func NewCfnWebhookPropsMixin(props *CfnWebhookMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWebhookPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWebhookPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWebhookPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnWebhookPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodePipeline::Webhook`.
func NewCfnWebhookPropsMixin_Override(c CfnWebhookPropsMixin, props *CfnWebhookMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnWebhookPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWebhookPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWebhookPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnWebhookPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebhookPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnWebhookPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWebhookPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWebhookPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

