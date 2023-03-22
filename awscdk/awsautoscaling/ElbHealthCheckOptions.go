package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// ELB Heath check options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elbHealthCheckOptions := &ElbHealthCheckOptions{
//   	Grace: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type ElbHealthCheckOptions struct {
	// Specified the time Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
	//
	// This option is required for ELB health checks.
	Grace awscdk.Duration `field:"required" json:"grace" yaml:"grace"`
}

