package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties for defining an Application Load Balancer.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
// Experimental.
type ApplicationLoadBalancerProps struct {
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
	// Indicates whether HTTP/2 is enabled.
	// Experimental.
	Http2Enabled *bool `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// The load balancer idle timeout, in seconds.
	// Experimental.
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// The type of IP addresses to use.
	//
	// Only applies to application load balancers.
	// Experimental.
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Security group to associate with this load balancer.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

