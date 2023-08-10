package awscdkappconfigalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   var application application
//
//
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
//   	Application: Application,
//   	Content: appconfig.ConfigurationContent_FromInline(jsii.String("This is my configuration content.")),
//   	DeploymentStrategy: appconfig.NewDeploymentStrategy(this, jsii.String("MyDeploymentStrategy"), &DeploymentStrategyProps{
//   		RolloutStrategy: appconfig.RolloutStrategy_Linear(&RolloutStrategyProps{
//   			GrowthFactor: jsii.Number(15),
//   			DeploymentDuration: awscdk.Duration_Minutes(jsii.Number(30)),
//   			FinalBakeTime: awscdk.Duration_*Minutes(jsii.Number(15)),
//   		}),
//   	}),
//   })
//
// Experimental.
type RolloutStrategyProps struct {
	// The deployment duration of the deployment strategy.
	//
	// This defines
	// the total amount of time for a deployment to last.
	// Experimental.
	DeploymentDuration awscdk.Duration `field:"required" json:"deploymentDuration" yaml:"deploymentDuration"`
	// The growth factor of the deployment strategy.
	//
	// This defines
	// the percentage of targets to receive a deployed configuration
	// during each interval.
	// Experimental.
	GrowthFactor *float64 `field:"required" json:"growthFactor" yaml:"growthFactor"`
	// The final bake time of the deployment strategy.
	//
	// This setting specifies the amount of time AWS AppConfig monitors for Amazon
	// CloudWatch alarms after the configuration has been deployed to
	// 100% of its targets, before considering the deployment to be complete.
	// If an alarm is triggered during this time, AWS AppConfig rolls back
	// the deployment.
	// Experimental.
	FinalBakeTime awscdk.Duration `field:"optional" json:"finalBakeTime" yaml:"finalBakeTime"`
}

