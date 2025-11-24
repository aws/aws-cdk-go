package mixinsawsmedialive

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsmedialive/mixinsawsmedialive/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::MediaLive::EventBridgeRuleTemplate Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBridgeRuleTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnEventBridgeRuleTemplatePropsMixin(&CfnEventBridgeRuleTemplateMixinProps{
//   	Description: jsii.String("description"),
//   	EventTargets: []interface{}{
//   		&EventBridgeRuleTemplateTargetProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   	EventType: jsii.String("eventType"),
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-eventbridgeruletemplate.html
//
type CfnEventBridgeRuleTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEventBridgeRuleTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEventBridgeRuleTemplatePropsMixin
type jsiiProxy_CfnEventBridgeRuleTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEventBridgeRuleTemplatePropsMixin) Props() *CfnEventBridgeRuleTemplateMixinProps {
	var returns *CfnEventBridgeRuleTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBridgeRuleTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaLive::EventBridgeRuleTemplate`.
func NewCfnEventBridgeRuleTemplatePropsMixin(props *CfnEventBridgeRuleTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEventBridgeRuleTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEventBridgeRuleTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventBridgeRuleTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnEventBridgeRuleTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaLive::EventBridgeRuleTemplate`.
func NewCfnEventBridgeRuleTemplatePropsMixin_Override(c CfnEventBridgeRuleTemplatePropsMixin, props *CfnEventBridgeRuleTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnEventBridgeRuleTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEventBridgeRuleTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventBridgeRuleTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnEventBridgeRuleTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventBridgeRuleTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnEventBridgeRuleTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventBridgeRuleTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEventBridgeRuleTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

