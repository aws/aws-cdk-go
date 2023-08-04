package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// EC2 Heath check options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ec2HealthCheckOptions := &Ec2HealthCheckOptions{
//   	Grace: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type Ec2HealthCheckOptions struct {
	// Specified the time Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
	// Default: Duration.seconds(0)
	//
	Grace awscdk.Duration `field:"optional" json:"grace" yaml:"grace"`
}

