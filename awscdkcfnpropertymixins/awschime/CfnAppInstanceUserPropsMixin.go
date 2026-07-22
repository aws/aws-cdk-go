package awschime

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awschime/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Chime::AppInstanceUser.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAppInstanceUserPropsMixin := awscdkcfnpropertymixins.Aws_chime.NewCfnAppInstanceUserPropsMixin(&CfnAppInstanceUserMixinProps{
//   	AppInstanceArn: jsii.String("appInstanceArn"),
//   	AppInstanceUserId: jsii.String("appInstanceUserId"),
//   	ExpirationSettings: &ExpirationSettingsProperty{
//   		ExpirationCriterion: jsii.String("expirationCriterion"),
//   		ExpirationDays: jsii.Number(123),
//   	},
//   	Metadata: jsii.String("metadata"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstanceuser.html
//
type CfnAppInstanceUserPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAppInstanceUserMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAppInstanceUserPropsMixin
type jsiiProxy_CfnAppInstanceUserPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAppInstanceUserPropsMixin) Props() *CfnAppInstanceUserMixinProps {
	var returns *CfnAppInstanceUserMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppInstanceUserPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Chime::AppInstanceUser`.
func NewCfnAppInstanceUserPropsMixin(props *CfnAppInstanceUserMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAppInstanceUserPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAppInstanceUserPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppInstanceUserPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnAppInstanceUserPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Chime::AppInstanceUser`.
func NewCfnAppInstanceUserPropsMixin_Override(c CfnAppInstanceUserPropsMixin, props *CfnAppInstanceUserMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnAppInstanceUserPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAppInstanceUserPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppInstanceUserPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnAppInstanceUserPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppInstanceUserPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnAppInstanceUserPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppInstanceUserPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAppInstanceUserPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

