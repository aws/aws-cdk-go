package awseksv2


// Launch template property specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateSpec := &LaunchTemplateSpec{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	Version: jsii.String("version"),
//   }
//
type LaunchTemplateSpec struct {
	// The Launch template ID.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The launch template version to be used (optional).
	// Default: - the default version of the launch template.
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

