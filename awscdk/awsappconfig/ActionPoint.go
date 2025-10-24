package awsappconfig


// Defines Extension action points.
//
// Example:
//   var fn Function
//
//
//   appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
//   	Actions: []Action{
//   		appconfig.NewAction(&ActionProps{
//   			ActionPoints: []ActionPoint{
//   				appconfig.ActionPoint_ON_DEPLOYMENT_START,
//   			},
//   			EventDestination: appconfig.NewLambdaDestination(fn),
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions-about.html#working-with-appconfig-extensions-how-it-works-step-2
//
type ActionPoint string

const (
	ActionPoint_PRE_CREATE_HOSTED_CONFIGURATION_VERSION ActionPoint = "PRE_CREATE_HOSTED_CONFIGURATION_VERSION"
	ActionPoint_PRE_START_DEPLOYMENT ActionPoint = "PRE_START_DEPLOYMENT"
	ActionPoint_ON_DEPLOYMENT_START ActionPoint = "ON_DEPLOYMENT_START"
	ActionPoint_ON_DEPLOYMENT_STEP ActionPoint = "ON_DEPLOYMENT_STEP"
	ActionPoint_ON_DEPLOYMENT_BAKING ActionPoint = "ON_DEPLOYMENT_BAKING"
	ActionPoint_ON_DEPLOYMENT_COMPLETE ActionPoint = "ON_DEPLOYMENT_COMPLETE"
	ActionPoint_ON_DEPLOYMENT_ROLLED_BACK ActionPoint = "ON_DEPLOYMENT_ROLLED_BACK"
	ActionPoint_AT_DEPLOYMENT_TICK ActionPoint = "AT_DEPLOYMENT_TICK"
)

