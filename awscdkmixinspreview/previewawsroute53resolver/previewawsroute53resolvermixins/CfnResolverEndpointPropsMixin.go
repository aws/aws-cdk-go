package previewawsroute53resolvermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53resolver/previewawsroute53resolvermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Resolver endpoint. There are two types of Resolver endpoints, inbound and outbound:.
//
// - An *inbound Resolver endpoint* forwards DNS queries to the DNS service for a VPC from your network.
// - An *outbound Resolver endpoint* forwards DNS queries from the DNS service for a VPC to your network.
//
// > - You cannot update `ResolverEndpointType` and `IpAddresses` in the same request.
// > - When you update a dual-stack IP address, you must update both IP addresses. You can’t update only an IPv4 or IPv6 and keep an existing IP address.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResolverEndpointPropsMixin := awscdkmixinspreview.Mixins.NewCfnResolverEndpointPropsMixin(&CfnResolverEndpointMixinProps{
//   	Direction: jsii.String("direction"),
//   	IpAddresses: []interface{}{
//   		&IpAddressRequestProperty{
//   			Ip: jsii.String("ip"),
//   			Ipv6: jsii.String("ipv6"),
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OutpostArn: jsii.String("outpostArn"),
//   	PreferredInstanceType: jsii.String("preferredInstanceType"),
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	ResolverEndpointType: jsii.String("resolverEndpointType"),
//   	RniEnhancedMetricsEnabled: jsii.Boolean(false),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetNameServerMetricsEnabled: jsii.Boolean(false),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html
//
type CfnResolverEndpointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResolverEndpointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResolverEndpointPropsMixin
type jsiiProxy_CfnResolverEndpointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResolverEndpointPropsMixin) Props() *CfnResolverEndpointMixinProps {
	var returns *CfnResolverEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53Resolver::ResolverEndpoint`.
func NewCfnResolverEndpointPropsMixin(props *CfnResolverEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResolverEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResolverEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResolverEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.mixins.CfnResolverEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53Resolver::ResolverEndpoint`.
func NewCfnResolverEndpointPropsMixin_Override(c CfnResolverEndpointPropsMixin, props *CfnResolverEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.mixins.CfnResolverEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResolverEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResolverEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53resolver.mixins.CfnResolverEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53resolver.mixins.CfnResolverEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnResolverEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

