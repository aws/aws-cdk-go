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
//
//   var vpc vpc
//
//
//   // Create the load balancer in a VPC. 'internetFacing' is 'false'
//   // by default, which creates an internal load balancer.
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//
//   // Add a listener and open up the load balancer's security group
//   // to the world.
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//
//   	// 'open: true' is the default, you can leave it out if you want. Set it
//   	// to 'false' and use `listener.connections` if you want to be selective
//   	// about who can access the load balancer.
//   	open: jsii.Boolean(true),
//   })
//
//   // Create an AutoScaling group and add it as a load balancing
//   // target to the listener.
//   listener.addTargets(jsii.String("ApplicationFleet"), &addApplicationTargetsProps{
//   	port: jsii.Number(8080),
//   	targets: []iApplicationLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
type ApplicationLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates whether deletion protection is enabled.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether the load balancer has an internet-routable address.
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
	// Name of the load balancer.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// Which subnets place the load balancer in.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Determines how the load balancer handles requests that might pose a security risk to your application.
	DesyncMitigationMode DesyncMitigationMode `field:"optional" json:"desyncMitigationMode" yaml:"desyncMitigationMode"`
	// Indicates whether HTTP headers with invalid header fields are removed by the load balancer (true) or routed to targets (false).
	DropInvalidHeaderFields *bool `field:"optional" json:"dropInvalidHeaderFields" yaml:"dropInvalidHeaderFields"`
	// Indicates whether HTTP/2 is enabled.
	Http2Enabled *bool `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// The load balancer idle timeout, in seconds.
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// The type of IP addresses to use.
	//
	// Only applies to application load balancers.
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Security group to associate with this load balancer.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

