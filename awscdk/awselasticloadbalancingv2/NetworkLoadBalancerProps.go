package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for a network load balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   })
//   listener := lb.AddListener(jsii.String("listener"), &BaseNetworkListenerProps{
//   	Port: jsii.Number(80),
//   })
//   listener.AddTargets(jsii.String("target"), &AddNetworkTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
//   	DefaultIntegration: awscdk.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type NetworkLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates whether cross-zone load balancing is enabled.
	// Default: - false for Network Load Balancers and true for Application Load Balancers.
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
}

