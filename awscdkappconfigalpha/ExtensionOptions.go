package awscdkappconfigalpha


// Options for the Extension construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//
//   var parameter parameter
//
//   extensionOptions := &ExtensionOptions{
//   	Description: jsii.String("description"),
//   	ExtensionName: jsii.String("extensionName"),
//   	LatestVersionNumber: jsii.Number(123),
//   	Parameters: []*parameter{
//   		parameter,
//   	},
//   }
//
// Experimental.
type ExtensionOptions struct {
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
}

