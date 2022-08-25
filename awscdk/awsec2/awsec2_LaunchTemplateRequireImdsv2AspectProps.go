package awsec2


// Properties for `LaunchTemplateRequireImdsv2Aspect`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateRequireImdsv2AspectProps := &launchTemplateRequireImdsv2AspectProps{
//   	suppressWarnings: jsii.Boolean(false),
//   }
//
type LaunchTemplateRequireImdsv2AspectProps struct {
	// Whether warning annotations from this Aspect should be suppressed or not.
	SuppressWarnings *bool `field:"optional" json:"suppressWarnings" yaml:"suppressWarnings"`
}

