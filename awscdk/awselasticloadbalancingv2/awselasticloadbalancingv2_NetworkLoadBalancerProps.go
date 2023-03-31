package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties for a network load balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   listener := lb.addListener(jsii.String("listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addNetworkTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdk.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
// Experimental.
type NetworkLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates whether deletion protection is enabled.
	// Experimental.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether the load balancer has an internet-routable address.
	// Experimental.
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
	// Name of the load balancer.
	// Experimental.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// Which subnets place the load balancer in.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Indicates whether cross-zone load balancing is enabled.
	// Experimental.
	CrossZoneEnabled *bool `field:"optional" json:"crossZoneEnabled" yaml:"crossZoneEnabled"`
}

