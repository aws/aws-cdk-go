package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLoadBalancerPropsMixin := awscdkcfnpropertymixins.Aws_elasticloadbalancingv2.NewCfnLoadBalancerPropsMixin(&CfnLoadBalancerMixinProps{
//   	EnableCapacityReservationProvisionStabilize: jsii.Boolean(false),
//   	EnablePrefixForIpv6SourceNat: jsii.String("enablePrefixForIpv6SourceNat"),
//   	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic: jsii.String("enforceSecurityGroupInboundRulesOnPrivateLinkTraffic"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Ipv4IpamPoolId: jsii.String("ipv4IpamPoolId"),
//   	LoadBalancerAttributes: []interface{}{
//   		&LoadBalancerAttributeProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MinimumLoadBalancerCapacity: &MinimumLoadBalancerCapacityProperty{
//   		CapacityUnits: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	Scheme: jsii.String("scheme"),
//   	SecurityGroups: []interface{}{
//   		jsii.String("securityGroups"),
//   	},
//   	SubnetMappings: []interface{}{
//   		&SubnetMappingProperty{
//   			AllocationId: jsii.String("allocationId"),
//   			IPv6Address: jsii.String("iPv6Address"),
//   			PrivateIPv4Address: jsii.String("privateIPv4Address"),
//   			SourceNatIpv6Prefix: jsii.String("sourceNatIpv6Prefix"),
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	Subnets: []interface{}{
//   		jsii.String("subnets"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html
//
type CfnLoadBalancerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLoadBalancerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLoadBalancerPropsMixin
type jsiiProxy_CfnLoadBalancerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLoadBalancerPropsMixin) Props() *CfnLoadBalancerMixinProps {
	var returns *CfnLoadBalancerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElasticLoadBalancingV2::LoadBalancer`.
func NewCfnLoadBalancerPropsMixin(props *CfnLoadBalancerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLoadBalancerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLoadBalancerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLoadBalancerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnLoadBalancerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElasticLoadBalancingV2::LoadBalancer`.
func NewCfnLoadBalancerPropsMixin_Override(c CfnLoadBalancerPropsMixin, props *CfnLoadBalancerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnLoadBalancerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLoadBalancerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLoadBalancerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnLoadBalancerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoadBalancerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnLoadBalancerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoadBalancerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLoadBalancerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

