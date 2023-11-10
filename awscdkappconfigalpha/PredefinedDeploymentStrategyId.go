package awscdkappconfigalpha


// Defines the deployment strategy ID's of AWS AppConfig predefined strategies.
// See: https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html
//
// Experimental.
type PredefinedDeploymentStrategyId string

const (
	// **AWS Recommended**.
	//
	// This strategy processes the deployment exponentially using a 10% growth factor over 20 minutes.
	// AWS AppConfig recommends using this strategy for production deployments because it aligns with AWS best practices
	// for configuration deployments.
	// Experimental.
	PredefinedDeploymentStrategyId_CANARY_10_PERCENT_20_MINUTES PredefinedDeploymentStrategyId = "CANARY_10_PERCENT_20_MINUTES"
	// **Testing/Demonstration**.
	//
	// This strategy deploys the configuration to half of all targets every 30 seconds for a
	// one-minute deployment. AWS AppConfig recommends using this strategy only for testing or demonstration purposes because
	// it has a short duration and bake time.
	// Experimental.
	PredefinedDeploymentStrategyId_LINEAR_50_PERCENT_EVERY_30_SECONDS PredefinedDeploymentStrategyId = "LINEAR_50_PERCENT_EVERY_30_SECONDS"
	// **AWS Recommended**.
	//
	// This strategy deploys the configuration to 20% of all targets every six minutes for a 30 minute deployment.
	// AWS AppConfig recommends using this strategy for production deployments because it aligns with AWS best practices
	// for configuration deployments.
	// Experimental.
	PredefinedDeploymentStrategyId_LINEAR_20_PERCENT_EVERY_6_MINUTES PredefinedDeploymentStrategyId = "LINEAR_20_PERCENT_EVERY_6_MINUTES"
	// **Quick**.
	//
	// This strategy deploys the configuration to all targets immediately.
	// Experimental.
	PredefinedDeploymentStrategyId_ALL_AT_ONCE PredefinedDeploymentStrategyId = "ALL_AT_ONCE"
)

