package awsroute53resolver

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Amazon Route 53 Resolver on an Outpost.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnOutpostResolverPropsMixin := awscdkcfnpropertymixins.Aws_route53resolver.NewCfnOutpostResolverPropsMixin(&CfnOutpostResolverMixinProps{
//   	InstanceCount: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	OutpostArn: jsii.String("outpostArn"),
//   	PreferredInstanceType: jsii.String("preferredInstanceType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-outpostresolver.html
//
type CfnOutpostResolverPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnOutpostResolverMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOutpostResolverPropsMixin
type jsiiProxy_CfnOutpostResolverPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnOutpostResolverPropsMixin) Props() *CfnOutpostResolverMixinProps {
	var returns *CfnOutpostResolverMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutpostResolverPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53Resolver::OutpostResolver`.
func NewCfnOutpostResolverPropsMixin(props *CfnOutpostResolverMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnOutpostResolverPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOutpostResolverPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOutpostResolverPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnOutpostResolverPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53Resolver::OutpostResolver`.
func NewCfnOutpostResolverPropsMixin_Override(c CfnOutpostResolverPropsMixin, props *CfnOutpostResolverMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnOutpostResolverPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnOutpostResolverPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOutpostResolverPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnOutpostResolverPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOutpostResolverPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnOutpostResolverPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOutpostResolverPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOutpostResolverPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

