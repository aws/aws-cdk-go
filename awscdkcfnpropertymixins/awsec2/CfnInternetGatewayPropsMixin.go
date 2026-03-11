package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Allocates an internet gateway for use with a VPC.
//
// After creating the Internet gateway, you then attach it to a VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnInternetGatewayPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnInternetGatewayPropsMixin(&CfnInternetGatewayMixinProps{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internetgateway.html
//
type CfnInternetGatewayPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnInternetGatewayMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInternetGatewayPropsMixin
type jsiiProxy_CfnInternetGatewayPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnInternetGatewayPropsMixin) Props() *CfnInternetGatewayMixinProps {
	var returns *CfnInternetGatewayMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInternetGatewayPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::InternetGateway`.
func NewCfnInternetGatewayPropsMixin(props *CfnInternetGatewayMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnInternetGatewayPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInternetGatewayPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInternetGatewayPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnInternetGatewayPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::InternetGateway`.
func NewCfnInternetGatewayPropsMixin_Override(c CfnInternetGatewayPropsMixin, props *CfnInternetGatewayMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnInternetGatewayPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnInternetGatewayPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInternetGatewayPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnInternetGatewayPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInternetGatewayPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnInternetGatewayPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInternetGatewayPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInternetGatewayPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

