package awscdkappconfigalpha


// Options for the Environment construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//
//   var monitor monitor
//
//   environmentOptions := &EnvironmentOptions{
//   	Description: jsii.String("description"),
//   	EnvironmentName: jsii.String("environmentName"),
//   	Monitors: []*monitor{
//   		monitor,
//   	},
//   }
//
// Experimental.
type EnvironmentOptions struct {
	// The description of the environment.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the environment.
	// Default: - A name is generated.
	//
	// Experimental.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// The monitors for the environment.
	// Default: - No monitors.
	//
	// Experimental.
	Monitors *[]Monitor `field:"optional" json:"monitors" yaml:"monitors"`
}

