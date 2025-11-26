package previewawsroute53mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53/previewawsroute53mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Route53::KeySigningKey` resource creates a new key-signing key (KSK) in a hosted zone.
//
// The hosted zone ID is passed as a parameter in the KSK properties. You can specify the properties of this KSK using the `Name` , `Status` , and `KeyManagementServiceArn` properties of the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKeySigningKeyPropsMixin := awscdkmixinspreview.Mixins.NewCfnKeySigningKeyPropsMixin(&CfnKeySigningKeyMixinProps{
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	KeyManagementServiceArn: jsii.String("keyManagementServiceArn"),
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-keysigningkey.html
//
type CfnKeySigningKeyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnKeySigningKeyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnKeySigningKeyPropsMixin
type jsiiProxy_CfnKeySigningKeyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnKeySigningKeyPropsMixin) Props() *CfnKeySigningKeyMixinProps {
	var returns *CfnKeySigningKeyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeySigningKeyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53::KeySigningKey`.
func NewCfnKeySigningKeyPropsMixin(props *CfnKeySigningKeyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnKeySigningKeyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnKeySigningKeyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnKeySigningKeyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnKeySigningKeyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53::KeySigningKey`.
func NewCfnKeySigningKeyPropsMixin_Override(c CfnKeySigningKeyPropsMixin, props *CfnKeySigningKeyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnKeySigningKeyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnKeySigningKeyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnKeySigningKeyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnKeySigningKeyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKeySigningKeyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnKeySigningKeyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnKeySigningKeyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnKeySigningKeyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

