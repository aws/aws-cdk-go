package awscdkappconfigalpha


// Attributes of an existing AWS AppConfig environment to import it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alarm alarm
//   var application application
//   var role role
//
//   environmentAttributes := &EnvironmentAttributes{
//   	Application: application,
//   	EnvironmentId: jsii.String("environmentId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Monitors: []monitor{
//   		&monitor{
//   			Alarm: alarm,
//
//   			// the properties below are optional
//   			AlarmRole: role,
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// Experimental.
type EnvironmentAttributes struct {
	// The application associated with the environment.
	// Experimental.
	Application IApplication `field:"required" json:"application" yaml:"application"`
	// The ID of the environment.
	// Experimental.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// The description of the environment.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The monitors for the environment.
	// Experimental.
	Monitors *[]*Monitor `field:"optional" json:"monitors" yaml:"monitors"`
	// The name of the environment.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

