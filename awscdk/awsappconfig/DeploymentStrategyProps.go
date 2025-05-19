package awsappconfig


// Properties for DeploymentStrategy.
//
// Example:
//   appconfig.NewDeploymentStrategy(this, jsii.String("MyDeploymentStrategy"), &DeploymentStrategyProps{
//   	RolloutStrategy: appconfig.RolloutStrategy_Linear(&RolloutStrategyProps{
//   		GrowthFactor: jsii.Number(20),
//   		DeploymentDuration: awscdk.Duration_Minutes(jsii.Number(30)),
//   		FinalBakeTime: awscdk.Duration_*Minutes(jsii.Number(30)),
//   	}),
//   })
//
type DeploymentStrategyProps struct {
	// The rollout strategy for the deployment strategy.
	//
	// You can use predefined deployment
	// strategies, such as RolloutStrategy.ALL_AT_ONCE, RolloutStrategy.LINEAR_50_PERCENT_EVERY_30_SECONDS,
	// or RolloutStrategy.CANARY_10_PERCENT_20_MINUTES.
	RolloutStrategy RolloutStrategy `field:"required" json:"rolloutStrategy" yaml:"rolloutStrategy"`
	// A name for the deployment strategy.
	// Default: - A name is generated.
	//
	DeploymentStrategyName *string `field:"optional" json:"deploymentStrategyName" yaml:"deploymentStrategyName"`
	// A description of the deployment strategy.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

