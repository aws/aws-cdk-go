package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties used to define TCP Based healthchecks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   tcpHealthCheckOptions := &TcpHealthCheckOptions{
//   	HealthyThreshold: jsii.Number(123),
//   	Interval: duration,
//   	Timeout: duration,
//   	UnhealthyThreshold: jsii.Number(123),
//   }
//
// Experimental.
type TcpHealthCheckOptions struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	// Experimental.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time period between each health check execution.
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The amount of time to wait when receiving a response from the health check.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive failed health checks that must occur before declaring a listener unhealthy.
	// Experimental.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

