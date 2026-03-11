package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The structure representing the address lists and address list attribute that will be used in evaluation of boolean expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMailManagerAddressListPropsMixin := awscdkcfnpropertymixins.Aws_ses.NewCfnMailManagerAddressListPropsMixin(&CfnMailManagerAddressListMixinProps{
//   	AddressListName: jsii.String("addressListName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageraddresslist.html
//
type CfnMailManagerAddressListPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMailManagerAddressListMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMailManagerAddressListPropsMixin
type jsiiProxy_CfnMailManagerAddressListPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMailManagerAddressListPropsMixin) Props() *CfnMailManagerAddressListMixinProps {
	var returns *CfnMailManagerAddressListMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMailManagerAddressListPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::MailManagerAddressList`.
func NewCfnMailManagerAddressListPropsMixin(props *CfnMailManagerAddressListMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMailManagerAddressListPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMailManagerAddressListPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMailManagerAddressListPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddressListPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::MailManagerAddressList`.
func NewCfnMailManagerAddressListPropsMixin_Override(c CfnMailManagerAddressListPropsMixin, props *CfnMailManagerAddressListMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddressListPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMailManagerAddressListPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMailManagerAddressListPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddressListPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMailManagerAddressListPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddressListPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMailManagerAddressListPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMailManagerAddressListPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

