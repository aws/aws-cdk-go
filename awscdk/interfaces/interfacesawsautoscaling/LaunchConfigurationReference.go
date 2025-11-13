package interfacesawsautoscaling


// A reference to a LaunchConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchConfigurationReference := &LaunchConfigurationReference{
//   	LaunchConfigurationName: jsii.String("launchConfigurationName"),
//   }
//
type LaunchConfigurationReference struct {
	// The LaunchConfigurationName of the LaunchConfiguration resource.
	LaunchConfigurationName *string `field:"required" json:"launchConfigurationName" yaml:"launchConfigurationName"`
}

