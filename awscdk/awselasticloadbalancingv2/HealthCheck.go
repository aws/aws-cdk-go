package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for configuring a health check.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	Cpu: jsii.Number(512),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		EntryPoint: []*string{
//   			jsii.String("entry"),
//   			jsii.String("point"),
//   		},
//   	},
//   })
//
//   loadBalancedFargateService.TargetGroup.ConfigureHealthCheck(&HealthCheck{
//   	Path: jsii.String("/custom-health-path"),
//   })
//
type HealthCheck struct {
	// Indicates whether health checks are enabled.
	//
	// If the target type is lambda,
	// health checks are disabled by default but can be enabled. If the target type
	// is instance or ip, health checks are always enabled and cannot be disabled.
	// Default: - Determined automatically.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// GRPC code to use when checking for a successful response from a target.
	//
	// You can specify values between 0 and 99. You can specify multiple values
	// (for example, "0,1") or a range of values (for example, "0-5").
	// Default: - 12.
	//
	HealthyGrpcCodes *string `field:"optional" json:"healthyGrpcCodes" yaml:"healthyGrpcCodes"`
	// HTTP code to use when checking for a successful response from a target.
	//
	// For Application Load Balancers, you can specify values between 200 and
	// 499, and the default value is 200. You can specify multiple values (for
	// example, "200,202") or a range of values (for example, "200-299").
	HealthyHttpCodes *string `field:"optional" json:"healthyHttpCodes" yaml:"healthyHttpCodes"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy.
	//
	// For Application Load Balancers, the default is 5. For Network Load Balancers, the default is 3.
	// Default: 5 for ALBs, 3 for NLBs.
	//
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// The approximate number of seconds between health checks for an individual target.
	//
	// Must be 5 to 300 seconds.
	// Default: 10 seconds if protocol is `GENEVE`, 35 seconds if target type is `lambda`, else 30 seconds.
	//
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The ping path destination where Elastic Load Balancing sends health check requests.
	// Default: /.
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port that the load balancer uses when performing health checks on the targets.
	// Default: 'traffic-port'.
	//
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol the load balancer uses when performing health checks on targets.
	//
	// The TCP protocol is supported for health checks only if the protocol of the target group is TCP, TLS, UDP, or TCP_UDP.
	// The TLS, UDP, and TCP_UDP protocols are not supported for health checks.
	// Default: HTTP for ALBs, TCP for NLBs.
	//
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	//
	// Must be 2 to 120 seconds.
	// Default: 6 seconds if the protocol is HTTP, 5 seconds if protocol is `GENEVE`, 30 seconds if target type is `lambda`, 10 seconds for TCP, TLS, or HTTPS.
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive health check failures required before considering a target unhealthy.
	//
	// For Application Load Balancers, the default is 2. For Network Load
	// Balancers, this value must be the same as the healthy threshold count.
	// Default: 2.
	//
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
}

