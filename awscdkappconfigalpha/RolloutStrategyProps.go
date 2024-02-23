package awscdkappconfigalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for the Rollout Strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rolloutStrategyProps := &RolloutStrategyProps{
//   	DeploymentDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   	GrowthFactor: jsii.Number(123),
//
//   	// the properties below are optional
//   	FinalBakeTime: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
// Deprecated.
type RolloutStrategyProps struct {
	// The deployment duration of the deployment strategy.
	//
	// This defines
	// the total amount of time for a deployment to last.
	// Deprecated.
	DeploymentDuration awscdk.Duration `field:"required" json:"deploymentDuration" yaml:"deploymentDuration"`
	// The growth factor of the deployment strategy.
	//
	// This defines
	// the percentage of targets to receive a deployed configuration
	// during each interval.
	// Deprecated.
	GrowthFactor *float64 `field:"required" json:"growthFactor" yaml:"growthFactor"`
	// The final bake time of the deployment strategy.
	//
	// This setting specifies the amount of time AWS AppConfig monitors for Amazon
	// CloudWatch alarms after the configuration has been deployed to
	// 100% of its targets, before considering the deployment to be complete.
	// If an alarm is triggered during this time, AWS AppConfig rolls back
	// the deployment.
	// Default: Duration.minutes(0)
	//
	// Deprecated.
	FinalBakeTime awscdk.Duration `field:"optional" json:"finalBakeTime" yaml:"finalBakeTime"`
}

