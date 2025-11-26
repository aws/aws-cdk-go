package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// [IPv6 only] Specifies an egress-only internet gateway for your VPC.
//
// An egress-only internet gateway is used to enable outbound communication over IPv6 from instances in your VPC to the internet, and prevents hosts outside of your VPC from initiating an IPv6 connection with your instance.
//
// For more information, see [Egress-only internet gateway](https://docs.aws.amazon.com/vpc/latest/userguide/egress-only-internet-gateway.html) in the *Amazon VPC User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEgressOnlyInternetGatewayPropsMixin := awscdkmixinspreview.Mixins.NewCfnEgressOnlyInternetGatewayPropsMixin(&CfnEgressOnlyInternetGatewayMixinProps{
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-egressonlyinternetgateway.html
//
type CfnEgressOnlyInternetGatewayPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEgressOnlyInternetGatewayMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEgressOnlyInternetGatewayPropsMixin
type jsiiProxy_CfnEgressOnlyInternetGatewayPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEgressOnlyInternetGatewayPropsMixin) Props() *CfnEgressOnlyInternetGatewayMixinProps {
	var returns *CfnEgressOnlyInternetGatewayMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEgressOnlyInternetGatewayPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::EgressOnlyInternetGateway`.
func NewCfnEgressOnlyInternetGatewayPropsMixin(props *CfnEgressOnlyInternetGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEgressOnlyInternetGatewayPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEgressOnlyInternetGatewayPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEgressOnlyInternetGatewayPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnEgressOnlyInternetGatewayPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::EgressOnlyInternetGateway`.
func NewCfnEgressOnlyInternetGatewayPropsMixin_Override(c CfnEgressOnlyInternetGatewayPropsMixin, props *CfnEgressOnlyInternetGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnEgressOnlyInternetGatewayPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEgressOnlyInternetGatewayPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEgressOnlyInternetGatewayPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnEgressOnlyInternetGatewayPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEgressOnlyInternetGatewayPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnEgressOnlyInternetGatewayPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEgressOnlyInternetGatewayPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEgressOnlyInternetGatewayPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

