package awscdkappconfigalpha


// Defines the deployment strategy ID's of AWS AppConfig predefined strategies.
// Experimental.
type PredefinedDeploymentStrategyId string

const (
	// Experimental.
	PredefinedDeploymentStrategyId_CANARY_10_PERCENT_20_MINUTES PredefinedDeploymentStrategyId = "CANARY_10_PERCENT_20_MINUTES"
	// Experimental.
	PredefinedDeploymentStrategyId_LINEAR_50_PERCENT_EVERY_30_SECONDS PredefinedDeploymentStrategyId = "LINEAR_50_PERCENT_EVERY_30_SECONDS"
	// Experimental.
	PredefinedDeploymentStrategyId_LINEAR_20_PERCENT_EVERY_6_MINUTES PredefinedDeploymentStrategyId = "LINEAR_20_PERCENT_EVERY_6_MINUTES"
	// Experimental.
	PredefinedDeploymentStrategyId_ALL_AT_ONCE PredefinedDeploymentStrategyId = "ALL_AT_ONCE"
)

