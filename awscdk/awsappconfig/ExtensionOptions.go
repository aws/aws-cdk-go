package awsappconfig


// Options for the Extension construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameter Parameter
//
//   extensionOptions := &ExtensionOptions{
//   	Description: jsii.String("description"),
//   	ExtensionName: jsii.String("extensionName"),
//   	LatestVersionNumber: jsii.Number(123),
//   	Parameters: []Parameter{
//   		parameter,
//   	},
//   }
//
type ExtensionOptions struct {
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
}

