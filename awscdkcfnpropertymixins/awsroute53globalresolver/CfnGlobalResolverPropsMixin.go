package awsroute53globalresolver

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53globalresolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::Route53GlobalResolver::GlobalResolver.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnGlobalResolverPropsMixin := awscdkcfnpropertymixins.Aws_route53globalresolver.NewCfnGlobalResolverPropsMixin(&CfnGlobalResolverMixinProps{
//   	ClientToken: jsii.String("clientToken"),
//   	Description: jsii.String("description"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Name: jsii.String("name"),
//   	ObservabilityRegion: jsii.String("observabilityRegion"),
//   	Regions: []*string{
//   		jsii.String("regions"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html
//
type CfnGlobalResolverPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnGlobalResolverMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGlobalResolverPropsMixin
type jsiiProxy_CfnGlobalResolverPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnGlobalResolverPropsMixin) Props() *CfnGlobalResolverMixinProps {
	var returns *CfnGlobalResolverMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalResolverPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53GlobalResolver::GlobalResolver`.
func NewCfnGlobalResolverPropsMixin(props *CfnGlobalResolverMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnGlobalResolverPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGlobalResolverPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGlobalResolverPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53globalresolver.CfnGlobalResolverPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53GlobalResolver::GlobalResolver`.
func NewCfnGlobalResolverPropsMixin_Override(c CfnGlobalResolverPropsMixin, props *CfnGlobalResolverMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53globalresolver.CfnGlobalResolverPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnGlobalResolverPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGlobalResolverPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53globalresolver.CfnGlobalResolverPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGlobalResolverPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53globalresolver.CfnGlobalResolverPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGlobalResolverPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGlobalResolverPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

