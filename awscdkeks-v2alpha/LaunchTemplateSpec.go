package awscdkeks-v2alpha


// Launch template property specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeks-v2alpha"
//
//   launchTemplateSpec := &LaunchTemplateSpec{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	Version: jsii.String("version"),
//   }
//
// Experimental.
type LaunchTemplateSpec struct {
	// The Launch template ID.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The launch template version to be used (optional).
	// Default: - the default version of the launch template.
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

