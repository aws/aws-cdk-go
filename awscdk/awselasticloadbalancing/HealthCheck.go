package awselasticloadbalancing

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Describe the health check to a load balancer.
//
// Example:
//   var vpc iVpc
//
//   var myAutoScalingGroup autoScalingGroup
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   	HealthCheck: &HealthCheck{
//   		Port: jsii.Number(80),
//   	},
//   })
//   lb.AddTarget(myAutoScalingGroup)
//   lb.AddListener(&LoadBalancerListener{
//   	ExternalPort: jsii.Number(80),
//   })
//
// Experimental.
type HealthCheck struct {
	// What port number to health check on.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// After how many successful checks is an instance considered healthy.
	// Experimental.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Number of seconds between health checks.
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// What path to use for HTTP or HTTPS health check (must return 200).
	//
	// For SSL and TCP health checks, accepting connections is enough to be considered
	// healthy.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// What protocol to use for health checking.
	//
	// The protocol is automatically determined from the port if it's not supplied.
	// Experimental.
	Protocol LoadBalancingProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Health check timeout.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// After how many unsuccessful checks is an instance considered unhealthy.
	// Experimental.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

