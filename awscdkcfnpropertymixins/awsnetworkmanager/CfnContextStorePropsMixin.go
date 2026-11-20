package awsnetworkmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Context Store paired to a Global Network for sharing network context with firewall vendors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnContextStorePropsMixin := awscdkcfnpropertymixins.Aws_networkmanager.NewCfnContextStorePropsMixin(&CfnContextStoreMixinProps{
//   	Description: jsii.String("description"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-contextstore.html
//
type CfnContextStorePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnContextStoreMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContextStorePropsMixin
type jsiiProxy_CfnContextStorePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnContextStorePropsMixin) Props() *CfnContextStoreMixinProps {
	var returns *CfnContextStoreMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContextStorePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkManager::ContextStore`.
func NewCfnContextStorePropsMixin(props *CfnContextStoreMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnContextStorePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContextStorePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContextStorePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnContextStorePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkManager::ContextStore`.
func NewCfnContextStorePropsMixin_Override(c CfnContextStorePropsMixin, props *CfnContextStoreMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnContextStorePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnContextStorePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContextStorePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnContextStorePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContextStorePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnContextStorePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContextStorePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnContextStorePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

