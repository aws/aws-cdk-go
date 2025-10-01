package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties used to define TCP Based healthchecks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tcpHealthCheckOptions := &TcpHealthCheckOptions{
//   	HealthyThreshold: jsii.Number(123),
//   	Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	UnhealthyThreshold: jsii.Number(123),
//   }
//
// Experimental.
type TcpHealthCheckOptions struct {
	// The number of consecutive checks that must succeed before App Runner decides that the service is healthy.
	// Default: 1.
	//
	// Experimental.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time interval, in seconds, between health checks.
	// Default: Duration.seconds(5)
	//
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The time, in seconds, to wait for a health check response before deciding it failed.
	// Default: Duration.seconds(2)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive checks that must fail before App Runner decides that the service is unhealthy.
	// Default: 5.
	//
	// Experimental.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

