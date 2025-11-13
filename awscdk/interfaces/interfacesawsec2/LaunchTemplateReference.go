package interfacesawsec2


// A reference to a LaunchTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateReference := &LaunchTemplateReference{
//   	LaunchTemplateId: jsii.String("launchTemplateId"),
//   }
//
type LaunchTemplateReference struct {
	// The LaunchTemplateId of the LaunchTemplate resource.
	LaunchTemplateId *string `field:"required" json:"launchTemplateId" yaml:"launchTemplateId"`
}

