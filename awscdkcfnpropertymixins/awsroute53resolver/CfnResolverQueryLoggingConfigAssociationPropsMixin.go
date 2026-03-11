package awsroute53resolver

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation resource is a configuration for DNS query logging.
//
// After you create a query logging configuration, Amazon Route 53 begins to publish log data to an Amazon CloudWatch Logs log group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnResolverQueryLoggingConfigAssociationPropsMixin := awscdkcfnpropertymixins.Aws_route53resolver.NewCfnResolverQueryLoggingConfigAssociationPropsMixin(&CfnResolverQueryLoggingConfigAssociationMixinProps{
//   	ResolverQueryLogConfigId: jsii.String("resolverQueryLogConfigId"),
//   	ResourceId: jsii.String("resourceId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverqueryloggingconfigassociation.html
//
type CfnResolverQueryLoggingConfigAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnResolverQueryLoggingConfigAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResolverQueryLoggingConfigAssociationPropsMixin
type jsiiProxy_CfnResolverQueryLoggingConfigAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociationPropsMixin) Props() *CfnResolverQueryLoggingConfigAssociationMixinProps {
	var returns *CfnResolverQueryLoggingConfigAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation`.
func NewCfnResolverQueryLoggingConfigAssociationPropsMixin(props *CfnResolverQueryLoggingConfigAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnResolverQueryLoggingConfigAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResolverQueryLoggingConfigAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResolverQueryLoggingConfigAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnResolverQueryLoggingConfigAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation`.
func NewCfnResolverQueryLoggingConfigAssociationPropsMixin_Override(c CfnResolverQueryLoggingConfigAssociationPropsMixin, props *CfnResolverQueryLoggingConfigAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnResolverQueryLoggingConfigAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnResolverQueryLoggingConfigAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResolverQueryLoggingConfigAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnResolverQueryLoggingConfigAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverQueryLoggingConfigAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53resolver.CfnResolverQueryLoggingConfigAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

