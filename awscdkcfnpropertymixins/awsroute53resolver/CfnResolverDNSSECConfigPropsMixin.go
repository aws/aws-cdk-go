package awsroute53resolver

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Route53Resolver::ResolverDNSSECConfig` resource is a complex type that contains information about a configuration for DNSSEC validation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnResolverDNSSECConfigPropsMixin := awscdkcfnpropertymixins.Aws_route53resolver.NewCfnResolverDNSSECConfigPropsMixin(&CfnResolverDNSSECConfigMixinProps{
//   	ResourceId: jsii.String("resourceId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverdnssecconfig.html
//
type CfnResolverDNSSECConfigPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnResolverDNSSECConfigMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResolverDNSSECConfigPropsMixin
type jsiiProxy_CfnResolverDNSSECConfigPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnResolverDNSSECConfigPropsMixin) Props() *CfnResolverDNSSECConfigMixinProps {
	var returns *CfnResolverDNSSECConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfigPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53Resolver::ResolverDNSSECConfig`.
func NewCfnResolverDNSSECConfigPropsMixin(props *CfnResolverDNSSECConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnResolverDNSSECConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResolverDNSSECConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResolverDNSSECConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnResolverDNSSECConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53Resolver::ResolverDNSSECConfig`.
func NewCfnResolverDNSSECConfigPropsMixin_Override(c CfnResolverDNSSECConfigPropsMixin, props *CfnResolverDNSSECConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnResolverDNSSECConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnResolverDNSSECConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResolverDNSSECConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnResolverDNSSECConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverDNSSECConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnResolverDNSSECConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverDNSSECConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

