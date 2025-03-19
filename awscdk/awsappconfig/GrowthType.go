package awsappconfig


// Defines the growth type of the deployment strategy.
type GrowthType string

const (
	// AWS AppConfig will process the deployment by increments of the growth factor evenly distributed over the deployment.
	GrowthType_LINEAR GrowthType = "LINEAR"
	// AWS AppConfig will process the deployment exponentially using the following formula: `G*(2^N)`.
	//
	// In this formula, `G` is the step percentage specified by the user and `N`
	// is the number of steps until the configuration is deployed to all targets.
	GrowthType_EXPONENTIAL GrowthType = "EXPONENTIAL"
)

