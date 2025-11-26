package previewawsiotwirelessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotwireless/previewawsiotwirelessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A partner account.
//
// If `PartnerAccountId` and `PartnerType` are `null` , returns all partner accounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPartnerAccountPropsMixin := awscdkmixinspreview.Mixins.NewCfnPartnerAccountPropsMixin(&CfnPartnerAccountMixinProps{
//   	AccountLinked: jsii.Boolean(false),
//   	PartnerAccountId: jsii.String("partnerAccountId"),
//   	PartnerType: jsii.String("partnerType"),
//   	Sidewalk: &SidewalkAccountInfoProperty{
//   		AppServerPrivateKey: jsii.String("appServerPrivateKey"),
//   	},
//   	SidewalkResponse: &SidewalkAccountInfoWithFingerprintProperty{
//   		AmazonId: jsii.String("amazonId"),
//   		Arn: jsii.String("arn"),
//   		Fingerprint: jsii.String("fingerprint"),
//   	},
//   	SidewalkUpdate: &SidewalkUpdateAccountProperty{
//   		AppServerPrivateKey: jsii.String("appServerPrivateKey"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html
//
type CfnPartnerAccountPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPartnerAccountMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPartnerAccountPropsMixin
type jsiiProxy_CfnPartnerAccountPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPartnerAccountPropsMixin) Props() *CfnPartnerAccountMixinProps {
	var returns *CfnPartnerAccountMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartnerAccountPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTWireless::PartnerAccount`.
func NewCfnPartnerAccountPropsMixin(props *CfnPartnerAccountMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPartnerAccountPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPartnerAccountPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPartnerAccountPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnPartnerAccountPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTWireless::PartnerAccount`.
func NewCfnPartnerAccountPropsMixin_Override(c CfnPartnerAccountPropsMixin, props *CfnPartnerAccountMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnPartnerAccountPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPartnerAccountPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPartnerAccountPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnPartnerAccountPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPartnerAccountPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnPartnerAccountPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPartnerAccountPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPartnerAccountPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

