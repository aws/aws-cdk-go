package awscdkappconfigalpha


// Defines Extension action points.
// See: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions-about.html#working-with-appconfig-extensions-how-it-works-step-2
//
// Deprecated.
type ActionPoint string

const (
	// Deprecated.
	ActionPoint_PRE_CREATE_HOSTED_CONFIGURATION_VERSION ActionPoint = "PRE_CREATE_HOSTED_CONFIGURATION_VERSION"
	// Deprecated.
	ActionPoint_PRE_START_DEPLOYMENT ActionPoint = "PRE_START_DEPLOYMENT"
	// Deprecated.
	ActionPoint_ON_DEPLOYMENT_START ActionPoint = "ON_DEPLOYMENT_START"
	// Deprecated.
	ActionPoint_ON_DEPLOYMENT_STEP ActionPoint = "ON_DEPLOYMENT_STEP"
	// Deprecated.
	ActionPoint_ON_DEPLOYMENT_BAKING ActionPoint = "ON_DEPLOYMENT_BAKING"
	// Deprecated.
	ActionPoint_ON_DEPLOYMENT_COMPLETE ActionPoint = "ON_DEPLOYMENT_COMPLETE"
	// Deprecated.
	ActionPoint_ON_DEPLOYMENT_ROLLED_BACK ActionPoint = "ON_DEPLOYMENT_ROLLED_BACK"
)

