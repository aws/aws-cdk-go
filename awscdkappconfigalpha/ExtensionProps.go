package awscdkappconfigalpha


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
// Experimental.
type ExtensionProps struct {
	// A description of the extension.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The latest version number of the extension.
	//
	// When you create a new version,
	// specify the most recent current version number. For example, you create version 3,
	// enter 2 for this field.
	// Experimental.
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// The name of the extension.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameters accepted for the extension.
	// Experimental.
	Parameters *[]Parameter `field:"optional" json:"parameters" yaml:"parameters"`
	// The actions for the extension.
	// Experimental.
	Actions *[]Action `field:"required" json:"actions" yaml:"actions"`
}

