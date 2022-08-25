package awsec2


// Properties for `InstanceRequireImdsv2Aspect`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceRequireImdsv2AspectProps := &instanceRequireImdsv2AspectProps{
//   	suppressLaunchTemplateWarning: jsii.Boolean(false),
//   	suppressWarnings: jsii.Boolean(false),
//   }
//
type InstanceRequireImdsv2AspectProps struct {
	// Whether warnings that would be raised when an Instance is associated with an existing Launch Template should be suppressed or not.
	//
	// You can set this to `true` if `LaunchTemplateImdsAspect` is being used alongside this Aspect to
	// suppress false-positive warnings because any Launch Templates associated with Instances will still be covered.
	SuppressLaunchTemplateWarning *bool `field:"optional" json:"suppressLaunchTemplateWarning" yaml:"suppressLaunchTemplateWarning"`
	// Whether warning annotations from this Aspect should be suppressed or not.
	SuppressWarnings *bool `field:"optional" json:"suppressWarnings" yaml:"suppressWarnings"`
}

