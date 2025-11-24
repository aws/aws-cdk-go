package mixinsawsssmcontacts

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsssmcontacts/mixinsawsssmcontacts/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SSMContacts::Contact` resource specifies a contact or escalation plan.
//
// Incident Manager contacts are a subset of actions and data types that you can use for managing responder engagement and interaction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContactPropsMixin := awscdkmixinspreview.Mixins.NewCfnContactPropsMixin(&CfnContactMixinProps{
//   	Alias: jsii.String("alias"),
//   	DisplayName: jsii.String("displayName"),
//   	Plan: []interface{}{
//   		&StageProperty{
//   			DurationInMinutes: jsii.Number(123),
//   			RotationIds: []*string{
//   				jsii.String("rotationIds"),
//   			},
//   			Targets: []interface{}{
//   				&TargetsProperty{
//   					ChannelTargetInfo: &ChannelTargetInfoProperty{
//   						ChannelId: jsii.String("channelId"),
//   						RetryIntervalInMinutes: jsii.Number(123),
//   					},
//   					ContactTargetInfo: &ContactTargetInfoProperty{
//   						ContactId: jsii.String("contactId"),
//   						IsEssential: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html
//
type CfnContactPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnContactMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContactPropsMixin
type jsiiProxy_CfnContactPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnContactPropsMixin) Props() *CfnContactMixinProps {
	var returns *CfnContactMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSMContacts::Contact`.
func NewCfnContactPropsMixin(props *CfnContactMixinProps, options *mixins.CfnPropertyMixinOptions) CfnContactPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContactPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContactPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmcontacts.mixins.CfnContactPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSMContacts::Contact`.
func NewCfnContactPropsMixin_Override(c CfnContactPropsMixin, props *CfnContactMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmcontacts.mixins.CfnContactPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnContactPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContactPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssmcontacts.mixins.CfnContactPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContactPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ssmcontacts.mixins.CfnContactPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContactPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnContactPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

