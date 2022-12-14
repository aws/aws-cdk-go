package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for adding new network targets to a listener.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
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
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type AddNetworkTargetsProps struct {
	// The port on which the listener listens for requests.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Indicates whether client IP preservation is enabled.
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Protocol for target group, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether Proxy Protocol version 2 is enabled.
	ProxyProtocolV2 *bool `field:"optional" json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	Targets *[]INetworkLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

