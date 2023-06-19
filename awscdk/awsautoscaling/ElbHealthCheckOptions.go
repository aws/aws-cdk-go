package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// ELB Heath check options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   elbHealthCheckOptions := &ElbHealthCheckOptions{
//   	Grace: duration,
//   }
//
// Experimental.
type ElbHealthCheckOptions struct {
	// Specified the time Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
	//
	// This option is required for ELB health checks.
	// Experimental.
	Grace awscdk.Duration `field:"required" json:"grace" yaml:"grace"`
}

