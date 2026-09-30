package interfacesawsmgn


// A reference to a LaunchConfigurationTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchConfigurationTemplateReference := &LaunchConfigurationTemplateReference{
//   	LaunchConfigurationTemplateArn: jsii.String("launchConfigurationTemplateArn"),
//   }
//
type LaunchConfigurationTemplateReference struct {
	// The Arn of the LaunchConfigurationTemplate resource.
	LaunchConfigurationTemplateArn *string `field:"required" json:"launchConfigurationTemplateArn" yaml:"launchConfigurationTemplateArn"`
}

