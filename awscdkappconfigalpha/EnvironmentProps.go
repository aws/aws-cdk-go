package awscdkappconfigalpha


// Properties for the Environment construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//
//   var application application
//   var monitor monitor
//
//   environmentProps := &EnvironmentProps{
//   	Application: application,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EnvironmentName: jsii.String("environmentName"),
//   	Monitors: []*monitor{
//   		monitor,
//   	},
//   }
//
// Deprecated.
type EnvironmentProps struct {
	// The description of the environment.
	// Default: - No description.
	//
	// Deprecated.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the environment.
	// Default: - A name is generated.
	//
	// Deprecated.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// The monitors for the environment.
	// Default: - No monitors.
	//
	// Deprecated.
	Monitors *[]Monitor `field:"optional" json:"monitors" yaml:"monitors"`
	// The application to be associated with the environment.
	// Deprecated.
	Application IApplication `field:"required" json:"application" yaml:"application"`
}

