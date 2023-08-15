package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties used to define GRPC Based healthchecks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcHealthCheckOptions := &GrpcHealthCheckOptions{
//   	HealthyThreshold: jsii.Number(123),
//   	Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	UnhealthyThreshold: jsii.Number(123),
//   }
//
type GrpcHealthCheckOptions struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	// Default: 2.
	//
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time period between each health check execution.
	// Default: Duration.seconds(5)
	//
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The amount of time to wait when receiving a response from the health check.
	// Default: Duration.seconds(2)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive failed health checks that must occur before declaring a listener unhealthy.
	// Default: - 2.
	//
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

