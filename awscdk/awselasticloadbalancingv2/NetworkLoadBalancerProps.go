package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for a network load balancer.
//
// Example:
//   var vpc Vpc
//
//
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
//   	EnablePrefixForIpv6SourceNat: jsii.Boolean(true),
//   })
//
//   listener := lb.AddListener(jsii.String("Listener"), &BaseNetworkListenerProps{
//   	Port: jsii.Number(1229),
//   	Protocol: elbv2.Protocol_UDP,
//   })
//
type NetworkLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates whether cross-zone load balancing is enabled.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattribute.html
	//
	// Default: - false for Network Load Balancers and true for Application Load Balancers.
	// This can not be `false` for Application Load Balancers.
	//
	CrossZoneEnabled *bool `field:"optional" json:"crossZoneEnabled" yaml:"crossZoneEnabled"`
	// Indicates whether deletion protection is enabled.
	// Default: false.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Indicates whether the load balancer blocks traffic through the Internet Gateway (IGW).
	// Default: - false for internet-facing load balancers and true for internal load balancers.
	//
	DenyAllIgwTraffic *bool `field:"optional" json:"denyAllIgwTraffic" yaml:"denyAllIgwTraffic"`
	// Whether the load balancer has an internet-routable address.
	// Default: false.
	//
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
	// Name of the load balancer.
	// Default: - Automatically generated name.
	//
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The minimum capacity (LCU) for a load balancer.
	// See: https://exampleloadbalancer.com/ondemand_capacity_reservation_calculator.html
	//
	// Default: undefined - ELB default is 0 LCU.
	//
	MinimumCapacityUnit *float64 `field:"optional" json:"minimumCapacityUnit" yaml:"minimumCapacityUnit"`
	// Which subnets place the load balancer in.
	// Default: - the Vpc default strategy.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The AZ affinity routing policy.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#zonal-dns-affinity
	//
	// Default: - AZ affinity is disabled.
	//
	ClientRoutingPolicy ClientRoutingPolicy `field:"optional" json:"clientRoutingPolicy" yaml:"clientRoutingPolicy"`
	// Create a Network Load Balancer without security groups.
	//
	// When true, creates an NLB that cannot have security groups attached.
	// This is useful when you need to create a traditional NLB without security group associations.
	//
	// This property only takes effect when the feature flag
	// `@aws-cdk/aws-elasticloadbalancingv2:networkLoadBalancerWithSecurityGroupByDefault` is enabled.
	// Default: false.
	//
	DisableSecurityGroups *bool `field:"optional" json:"disableSecurityGroups" yaml:"disableSecurityGroups"`
	// Indicates whether to use an IPv6 prefix from each subnet for source NAT.
	//
	// The IP address type must be IpAddressType.DUALSTACK.
	// Default: undefined - NLB default behavior is false.
	//
	EnablePrefixForIpv6SourceNat *bool `field:"optional" json:"enablePrefixForIpv6SourceNat" yaml:"enablePrefixForIpv6SourceNat"`
	// Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through AWS PrivateLink.
	// Default: true.
	//
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic *bool `field:"optional" json:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic" yaml:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic"`
	// The type of IP addresses to use.
	//
	// If you want to add a UDP or TCP_UDP listener to the load balancer,
	// you must choose IPv4.
	// Default: IpAddressType.IPV4
	//
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Security groups to associate with this load balancer.
	// Default: - No security groups associated with the load balancer.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Subnet information for the load balancer.
	// Default: undefined - The VPC default strategy for subnets is used.
	//
	SubnetMappings *[]*SubnetMapping `field:"optional" json:"subnetMappings" yaml:"subnetMappings"`
	// Indicates whether zonal shift is enabled.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/zonal-shift.html
	//
	// Default: false.
	//
	ZonalShift *bool `field:"optional" json:"zonalShift" yaml:"zonalShift"`
}

