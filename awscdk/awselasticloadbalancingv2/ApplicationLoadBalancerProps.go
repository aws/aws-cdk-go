package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for defining an Application Load Balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("lb"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   })
//   listener := lb.AddListener(jsii.String("listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   listener.AddTargets(jsii.String("target"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
//   	DefaultIntegration: awscdk.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type ApplicationLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates whether deletion protection is enabled.
	// Default: false.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
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
	// Determines how the load balancer handles requests that might pose a security risk to your application.
	// Default: DesyncMitigationMode.DEFENSIVE
	//
	DesyncMitigationMode DesyncMitigationMode `field:"optional" json:"desyncMitigationMode" yaml:"desyncMitigationMode"`
	// Indicates whether HTTP headers with invalid header fields are removed by the load balancer (true) or routed to targets (false).
	// Default: false.
	//
	DropInvalidHeaderFields *bool `field:"optional" json:"dropInvalidHeaderFields" yaml:"dropInvalidHeaderFields"`
	// Indicates whether HTTP/2 is enabled.
	// Default: true.
	//
	Http2Enabled *bool `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// The load balancer idle timeout, in seconds.
	// Default: 60.
	//
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// The type of IP addresses to use.
	//
	// Only applies to application load balancers.
	// Default: IpAddressType.Ipv4
	//
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Security group to associate with this load balancer.
	// Default: A security group is created.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

