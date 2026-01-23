package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The prefix to use for source NAT for a dual-stack network load balancer with UDP listeners.
//
// Example:
//   var vpc IVpc
//   var dualstackVpc IVpc
//   var subnet ISubnet
//   var dualstackSubnet ISubnet
//   var cfnEip CfnEIP
//
//
//   // Internet facing Network Load Balancer with an Elastic IPv4 address
//   // Internet facing Network Load Balancer with an Elastic IPv4 address
//   elbv2.NewNetworkLoadBalancer(this, jsii.String("InternetFacingLb"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   	SubnetMappings: []SubnetMapping{
//   		&SubnetMapping{
//   			Subnet: *Subnet,
//   			// The allocation ID of the Elastic IP address
//   			AllocationId: cfnEip.attrAllocationId,
//   		},
//   	},
//   })
//
//   // Internal Network Load Balancer with a private IPv4 address
//   // Internal Network Load Balancer with a private IPv4 address
//   elbv2.NewNetworkLoadBalancer(this, jsii.String("InternalLb"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(false),
//   	SubnetMappings: []SubnetMapping{
//   		&SubnetMapping{
//   			Subnet: *Subnet,
//   			// The private IPv4 address from the subnet
//   			// The address must be in the subnet's CIDR range and
//   			// can not be configured for a internet facing Network Load Balancer.
//   			PrivateIpv4Address: jsii.String("10.0.12.29"),
//   		},
//   	},
//   })
//
//   // Dualstack Network Load Balancer with an IPv6 address and prefix for source NAT
//   // Dualstack Network Load Balancer with an IPv6 address and prefix for source NAT
//   elbv2.NewNetworkLoadBalancer(this, jsii.String("DualstackLb"), &NetworkLoadBalancerProps{
//   	Vpc: dualstackVpc,
//   	// Configure the dualstack Network Load Balancer
//   	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
//   	EnablePrefixForIpv6SourceNat: jsii.Boolean(true),
//   	SubnetMappings: []SubnetMapping{
//   		&SubnetMapping{
//   			Subnet: dualstackSubnet,
//   			// The IPv6 address from the subnet
//   			// `ipAddresstype` must be `DUAL_STACK` or `DUAL_STACK_WITHOUT_PUBLIC_IPV4` to set the IPv6 address.
//   			Ipv6Address: jsii.String("2001:db8:1234:1a00::10"),
//   			// The IPv6 prefix to use for source NAT
//   			// `enablePrefixForIpv6SourceNat` must be `true` to set `sourceNatIpv6Prefix`.
//   			SourceNatIpv6Prefix: elbv2.SourceNatIpv6Prefix_AutoAssigned(),
//   		},
//   	},
//   })
//
type SourceNatIpv6Prefix interface {
	// The IPv6 prefix.
	Prefix() *string
}

// The jsii proxy struct for SourceNatIpv6Prefix
type jsiiProxy_SourceNatIpv6Prefix struct {
	_ byte // padding
}

func (j *jsiiProxy_SourceNatIpv6Prefix) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}


func NewSourceNatIpv6Prefix(prefix *string) SourceNatIpv6Prefix {
	_init_.Initialize()

	if err := validateNewSourceNatIpv6PrefixParameters(prefix); err != nil {
		panic(err)
	}
	j := jsiiProxy_SourceNatIpv6Prefix{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.SourceNatIpv6Prefix",
		[]interface{}{prefix},
		&j,
	)

	return &j
}

func NewSourceNatIpv6Prefix_Override(s SourceNatIpv6Prefix, prefix *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.SourceNatIpv6Prefix",
		[]interface{}{prefix},
		s,
	)
}

// Use an automatically assigned IPv6 prefix.
func SourceNatIpv6Prefix_AutoAssigned() SourceNatIpv6Prefix {
	_init_.Initialize()

	var returns SourceNatIpv6Prefix

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.SourceNatIpv6Prefix",
		"autoAssigned",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Use a custom IPv6 prefix with /80 netmask.
func SourceNatIpv6Prefix_FromIpv6Prefix(prefix *string) SourceNatIpv6Prefix {
	_init_.Initialize()

	if err := validateSourceNatIpv6Prefix_FromIpv6PrefixParameters(prefix); err != nil {
		panic(err)
	}
	var returns SourceNatIpv6Prefix

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.SourceNatIpv6Prefix",
		"fromIpv6Prefix",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

