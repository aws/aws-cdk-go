package awsiotwireless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A partner account.
//
// If `PartnerAccountId` and `PartnerType` are `null` , returns all partner accounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPartnerAccountPropsMixin := awscdkcfnpropertymixins.Aws_iotwireless.NewCfnPartnerAccountPropsMixin(&CfnPartnerAccountMixinProps{
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html
//
type CfnPartnerAccountPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPartnerAccountMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPartnerAccountPropsMixin
type jsiiProxy_CfnPartnerAccountPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnPartnerAccountPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTWireless::PartnerAccount`.
func NewCfnPartnerAccountPropsMixin(props *CfnPartnerAccountMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPartnerAccountPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPartnerAccountPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPartnerAccountPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnPartnerAccountPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTWireless::PartnerAccount`.
func NewCfnPartnerAccountPropsMixin_Override(c CfnPartnerAccountPropsMixin, props *CfnPartnerAccountMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnPartnerAccountPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPartnerAccountPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPartnerAccountPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnPartnerAccountPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnPartnerAccountPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPartnerAccountPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

