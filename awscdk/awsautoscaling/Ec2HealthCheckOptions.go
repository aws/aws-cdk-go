package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// EC2 Heath check options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   ec2HealthCheckOptions := &Ec2HealthCheckOptions{
//   	Grace: duration,
//   }
//
// Experimental.
type Ec2HealthCheckOptions struct {
	// Specified the time Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
	// Experimental.
	Grace awscdk.Duration `field:"optional" json:"grace" yaml:"grace"`
}

