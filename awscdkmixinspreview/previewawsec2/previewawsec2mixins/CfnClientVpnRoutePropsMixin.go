package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a network route to add to a Client VPN endpoint.
//
// Each Client VPN endpoint has a route table that describes the available destination network routes. Each route in the route table specifies the path for traffic to specific resources or networks.
//
// A target network association must be created before you can specify a route. If you're setting up all the components of a Client VPN endpoint at the same time, you must use the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the `AWS::EC2::ClientVpnTargetNetworkAssociation` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClientVpnRoutePropsMixin := awscdkmixinspreview.Mixins.NewCfnClientVpnRoutePropsMixin(&CfnClientVpnRouteMixinProps{
//   	ClientVpnEndpointId: jsii.String("clientVpnEndpointId"),
//   	Description: jsii.String("description"),
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	TargetVpcSubnetId: jsii.String("targetVpcSubnetId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html
//
type CfnClientVpnRoutePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnClientVpnRouteMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClientVpnRoutePropsMixin
type jsiiProxy_CfnClientVpnRoutePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnClientVpnRoutePropsMixin) Props() *CfnClientVpnRouteMixinProps {
	var returns *CfnClientVpnRouteMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnRoutePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::ClientVpnRoute`.
func NewCfnClientVpnRoutePropsMixin(props *CfnClientVpnRouteMixinProps, options *mixins.CfnPropertyMixinOptions) CfnClientVpnRoutePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClientVpnRoutePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClientVpnRoutePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnClientVpnRoutePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::ClientVpnRoute`.
func NewCfnClientVpnRoutePropsMixin_Override(c CfnClientVpnRoutePropsMixin, props *CfnClientVpnRouteMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnClientVpnRoutePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnClientVpnRoutePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnRoutePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnClientVpnRoutePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClientVpnRoutePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnClientVpnRoutePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClientVpnRoutePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnClientVpnRoutePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

