package awsnetworkmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDevicePropsMixin := awscdkcfnpropertymixins.Aws_networkmanager.NewCfnDevicePropsMixin(&CfnDeviceMixinProps{
//   	AwsLocation: &AWSLocationProperty{
//   		SubnetArn: jsii.String("subnetArn"),
//   		Zone: jsii.String("zone"),
//   	},
//   	Description: jsii.String("description"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	Location: &LocationProperty{
//   		Address: jsii.String("address"),
//   		Latitude: jsii.String("latitude"),
//   		Longitude: jsii.String("longitude"),
//   	},
//   	Model: jsii.String("model"),
//   	SerialNumber: jsii.String("serialNumber"),
//   	SiteId: jsii.String("siteId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	Vendor: jsii.String("vendor"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html
//
type CfnDevicePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDeviceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDevicePropsMixin
type jsiiProxy_CfnDevicePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDevicePropsMixin) Props() *CfnDeviceMixinProps {
	var returns *CfnDeviceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevicePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkManager::Device`.
func NewCfnDevicePropsMixin(props *CfnDeviceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDevicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDevicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDevicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDevicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkManager::Device`.
func NewCfnDevicePropsMixin_Override(c CfnDevicePropsMixin, props *CfnDeviceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDevicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDevicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDevicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDevicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDevicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDevicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDevicePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDevicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

