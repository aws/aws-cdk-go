package awscdkappconfigalpha


// Properties for DeploymentStrategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//
//   var rolloutStrategy rolloutStrategy
//
//   deploymentStrategyProps := &DeploymentStrategyProps{
//   	RolloutStrategy: rolloutStrategy,
//
//   	// the properties below are optional
//   	DeploymentStrategyName: jsii.String("deploymentStrategyName"),
//   	Description: jsii.String("description"),
//   }
//
// Deprecated.
type DeploymentStrategyProps struct {
	// The rollout strategy for the deployment strategy.
	//
	// You can use predefined deployment
	// strategies, such as RolloutStrategy.ALL_AT_ONCE, RolloutStrategy.LINEAR_50_PERCENT_EVERY_30_SECONDS,
	// or RolloutStrategy.CANARY_10_PERCENT_20_MINUTES.
	// Deprecated.
	RolloutStrategy RolloutStrategy `field:"required" json:"rolloutStrategy" yaml:"rolloutStrategy"`
	// A name for the deployment strategy.
	// Default: - A name is generated.
	//
	// Deprecated.
	DeploymentStrategyName *string `field:"optional" json:"deploymentStrategyName" yaml:"deploymentStrategyName"`
	// A description of the deployment strategy.
	// Default: - No description.
	//
	// Deprecated.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

