package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Cognito::UserPoolSessionConfigurationAttachment.
//
// Manages the session configuration of a user pool: whether user session tracking is enabled and the maximum number of concurrent token sessions per user. The service default (UserSessionTracking DISABLED with TokenSession MaxPerUser 20) cannot be declared; remove the resource to return the user pool to the default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnUserPoolSessionConfigurationAttachmentPropsMixin := awscdkcfnpropertymixins.Aws_cognito.NewCfnUserPoolSessionConfigurationAttachmentPropsMixin(&CfnUserPoolSessionConfigurationAttachmentMixinProps{
//   	TokenSession: &TokenSessionProperty{
//   		MaxPerUser: jsii.Number(123),
//   	},
//   	UserPoolId: jsii.String("userPoolId"),
//   	UserSessionTracking: jsii.String("userSessionTracking"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolsessionconfigurationattachment.html
//
type CfnUserPoolSessionConfigurationAttachmentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUserPoolSessionConfigurationAttachmentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserPoolSessionConfigurationAttachmentPropsMixin
type jsiiProxy_CfnUserPoolSessionConfigurationAttachmentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnUserPoolSessionConfigurationAttachmentPropsMixin) Props() *CfnUserPoolSessionConfigurationAttachmentMixinProps {
	var returns *CfnUserPoolSessionConfigurationAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolSessionConfigurationAttachmentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::UserPoolSessionConfigurationAttachment`.
func NewCfnUserPoolSessionConfigurationAttachmentPropsMixin(props *CfnUserPoolSessionConfigurationAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUserPoolSessionConfigurationAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPoolSessionConfigurationAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPoolSessionConfigurationAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolSessionConfigurationAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::UserPoolSessionConfigurationAttachment`.
func NewCfnUserPoolSessionConfigurationAttachmentPropsMixin_Override(c CfnUserPoolSessionConfigurationAttachmentPropsMixin, props *CfnUserPoolSessionConfigurationAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolSessionConfigurationAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUserPoolSessionConfigurationAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPoolSessionConfigurationAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolSessionConfigurationAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolSessionConfigurationAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolSessionConfigurationAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolSessionConfigurationAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUserPoolSessionConfigurationAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

