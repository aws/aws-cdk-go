package awscdkappconfigalpha


// Attributes of an existing AWS AppConfig environment to import it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//
//   var application application
//   var monitor monitor
//
//   environmentAttributes := &EnvironmentAttributes{
//   	Application: application,
//   	EnvironmentId: jsii.String("environmentId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Monitors: []*monitor{
//   		monitor,
//   	},
//   	Name: jsii.String("name"),
//   }
//
// Deprecated.
type EnvironmentAttributes struct {
	// The application associated with the environment.
	// Deprecated.
	Application IApplication `field:"required" json:"application" yaml:"application"`
	// The ID of the environment.
	// Deprecated.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// The description of the environment.
	// Default: - None.
	//
	// Deprecated.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The monitors for the environment.
	// Default: - None.
	//
	// Deprecated.
	Monitors *[]Monitor `field:"optional" json:"monitors" yaml:"monitors"`
	// The name of the environment.
	// Default: - None.
	//
	// Deprecated.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

