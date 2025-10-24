package awsappconfig


// Options for the Environment construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var monitor Monitor
//
//   environmentOptions := &EnvironmentOptions{
//   	DeletionProtectionCheck: awscdk.Aws_appconfig.DeletionProtectionCheck_ACCOUNT_DEFAULT,
//   	Description: jsii.String("description"),
//   	EnvironmentName: jsii.String("environmentName"),
//   	Monitors: []Monitor{
//   		monitor,
//   	},
//   }
//
type EnvironmentOptions struct {
	// A property to prevent accidental deletion of active environments.
	// Default: undefined - AppConfig default is ACCOUNT_DEFAULT.
	//
	DeletionProtectionCheck DeletionProtectionCheck `field:"optional" json:"deletionProtectionCheck" yaml:"deletionProtectionCheck"`
	// The description of the environment.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the environment.
	// Default: - A name is generated.
	//
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// The monitors for the environment.
	// Default: - No monitors.
	//
	Monitors *[]Monitor `field:"optional" json:"monitors" yaml:"monitors"`
}

