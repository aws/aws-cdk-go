package awsappconfig


// Properties for the Extension construct.
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
type ExtensionProps struct {
	// A description of the extension.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the extension.
	// Default: - A name is generated.
	//
	ExtensionName *string `field:"optional" json:"extensionName" yaml:"extensionName"`
	// The latest version number of the extension.
	//
	// When you create a new version,
	// specify the most recent current version number. For example, you create version 3,
	// enter 2 for this field.
	// Default: - None.
	//
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// The parameters accepted for the extension.
	// Default: - None.
	//
	Parameters *[]Parameter `field:"optional" json:"parameters" yaml:"parameters"`
	// The actions for the extension.
	Actions *[]Action `field:"required" json:"actions" yaml:"actions"`
}

