package awscdkappconfigalpha


// Example:
//   appconfig.NewDeploymentStrategy(this, jsii.String("MyDeploymentStrategy"), &DeploymentStrategyProps{
//   	RolloutStrategy: appconfig.RolloutStrategy_Linear(&RolloutStrategyProps{
//   		GrowthFactor: jsii.Number(20),
//   		DeploymentDuration: awscdk.Duration_Minutes(jsii.Number(30)),
//   		FinalBakeTime: awscdk.Duration_*Minutes(jsii.Number(30)),
//   	}),
//   })
//
// Experimental.
type DeploymentStrategyProps struct {
	// The rollout strategy for the deployment strategy.
	//
	// You can use predefined deployment
	// strategies, such as RolloutStrategy.ALL_AT_ONCE, RolloutStrategy.LINEAR_50_PERCENT_EVERY_30_SECONDS,
	// or RolloutStrategy.CANARY_10_PERCENT_20_MINUTES.
	// Experimental.
	RolloutStrategy RolloutStrategy `field:"required" json:"rolloutStrategy" yaml:"rolloutStrategy"`
	// A description of the deployment strategy.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the deployment strategy.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

