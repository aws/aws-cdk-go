package awscdkappconfigalpha


// Properties for the Extension construct.
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
// Experimental.
type ExtensionProps struct {
	// A description of the extension.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the extension.
	// Default: - A name is generated.
	//
	// Experimental.
	ExtensionName *string `field:"optional" json:"extensionName" yaml:"extensionName"`
	// The latest version number of the extension.
	//
	// When you create a new version,
	// specify the most recent current version number. For example, you create version 3,
	// enter 2 for this field.
	// Default: - None.
	//
	// Experimental.
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// The parameters accepted for the extension.
	// Default: - None.
	//
	// Experimental.
	Parameters *[]Parameter `field:"optional" json:"parameters" yaml:"parameters"`
	// The actions for the extension.
	// Experimental.
	Actions *[]Action `field:"required" json:"actions" yaml:"actions"`
}

