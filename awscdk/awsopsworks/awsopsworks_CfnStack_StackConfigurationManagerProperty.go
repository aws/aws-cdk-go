package awsopsworks


// Describes the configuration manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackConfigurationManagerProperty := &stackConfigurationManagerProperty{
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnStack_StackConfigurationManagerProperty struct {
	// The name.
	//
	// This parameter must be set to `Chef` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Chef version.
	//
	// This parameter must be set to 12, 11.10, or 11.4 for Linux stacks, and to 12.2 for Windows stacks. The default value for Linux stacks is 12.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

