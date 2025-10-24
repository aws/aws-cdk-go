package awsappconfig


// Attributes of an existing AWS AppConfig environment to import it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var application Application
//   var monitor Monitor
//
//   environmentAttributes := &EnvironmentAttributes{
//   	Application: application,
//   	EnvironmentId: jsii.String("environmentId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Monitors: []Monitor{
//   		monitor,
//   	},
//   	Name: jsii.String("name"),
//   }
//
type EnvironmentAttributes struct {
	// The application associated with the environment.
	Application IApplication `field:"required" json:"application" yaml:"application"`
	// The ID of the environment.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// The description of the environment.
	// Default: - None.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The monitors for the environment.
	// Default: - None.
	//
	Monitors *[]Monitor `field:"optional" json:"monitors" yaml:"monitors"`
	// The name of the environment.
	// Default: - None.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

