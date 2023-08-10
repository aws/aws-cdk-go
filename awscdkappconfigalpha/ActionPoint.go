package awscdkappconfigalpha


// Defines Extension action points.
//
// Example:
//   var fn function
//
//
//   appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
//   	Actions: []action{
//   		appconfig.NewAction(&ActionProps{
//   			ActionPoints: []actionPoint{
//   				appconfig.*actionPoint_ON_DEPLOYMENT_START,
//   			},
//   			EventDestination: appconfig.NewLambdaDestination(fn),
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions-about.html#working-with-appconfig-extensions-how-it-works-step-2
//
// Experimental.
type ActionPoint string

const (
	// Experimental.
	ActionPoint_PRE_CREATE_HOSTED_CONFIGURATION_VERSION ActionPoint = "PRE_CREATE_HOSTED_CONFIGURATION_VERSION"
	// Experimental.
	ActionPoint_PRE_START_DEPLOYMENT ActionPoint = "PRE_START_DEPLOYMENT"
	// Experimental.
	ActionPoint_ON_DEPLOYMENT_START ActionPoint = "ON_DEPLOYMENT_START"
	// Experimental.
	ActionPoint_ON_DEPLOYMENT_STEP ActionPoint = "ON_DEPLOYMENT_STEP"
	// Experimental.
	ActionPoint_ON_DEPLOYMENT_BAKING ActionPoint = "ON_DEPLOYMENT_BAKING"
	// Experimental.
	ActionPoint_ON_DEPLOYMENT_COMPLETE ActionPoint = "ON_DEPLOYMENT_COMPLETE"
	// Experimental.
	ActionPoint_ON_DEPLOYMENT_ROLLED_BACK ActionPoint = "ON_DEPLOYMENT_ROLLED_BACK"
)

