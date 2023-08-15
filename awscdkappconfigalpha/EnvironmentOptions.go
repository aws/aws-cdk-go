package awscdkappconfigalpha


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alarm alarm
//   var role role
//
//   environmentOptions := &EnvironmentOptions{
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
type EnvironmentOptions struct {
	// The description of the environment.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The monitors for the environment.
	// Default: - No monitors.
	//
	// Experimental.
	Monitors *[]*Monitor `field:"optional" json:"monitors" yaml:"monitors"`
	// The name of the environment.
	// Default: - A name is generated.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

