package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for defining an Application Load Balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var asg autoScalingGroup
//   var vpc vpc
//
//
//   // Create the load balancer in a VPC. 'internetFacing' is 'false'
//   // by default, which creates an internal load balancer.
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   // Add a listener and open up the load balancer's security group
//   // to the world.
//   listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//
//   	// 'open: true' is the default, you can leave it out if you want. Set it
//   	// to 'false' and use `listener.connections` if you want to be selective
//   	// about who can access the load balancer.
//   	Open: jsii.Boolean(true),
//   })
//
//   // Create an AutoScaling group and add it as a load balancing
//   // target to the listener.
//   listener.AddTargets(jsii.String("ApplicationFleet"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(8080),
//   	Targets: []iApplicationLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
type ApplicationLoadBalancerProps struct {
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
	// The client keep alive duration.
	//
	// The valid range is 60 to 604800 seconds (1 minute to 7 days).
	// Default: - Duration.seconds(3600)
	//
	ClientKeepAlive awscdk.Duration `field:"optional" json:"clientKeepAlive" yaml:"clientKeepAlive"`
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
	// Default: IpAddressType.IPV4
	//
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Security group to associate with this load balancer.
	// Default: A security group is created.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

