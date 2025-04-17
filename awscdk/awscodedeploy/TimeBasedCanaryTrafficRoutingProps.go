package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for `TimeBasedCanaryTrafficRouting`.
//
// Example:
//   config := codedeploy.NewLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &LambdaDeploymentConfigProps{
//   	TrafficRouting: codedeploy.NewTimeBasedCanaryTrafficRouting(&TimeBasedCanaryTrafficRoutingProps{
//   		Interval: awscdk.Duration_Minutes(jsii.Number(15)),
//   		Percentage: jsii.Number(5),
//   	}),
//   	DeploymentConfigName: jsii.String("MyDeploymentConfig"),
//   })
//
type TimeBasedCanaryTrafficRoutingProps struct {
	// The amount of time between traffic shifts.
	Interval awscdk.Duration `field:"required" json:"interval" yaml:"interval"`
	// The percentage to increase traffic on each traffic shift.
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

