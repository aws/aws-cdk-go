package awsimagebuilderalpha


// Properties for an EC2 Image Builder Amazon-managed component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   amazonManagedComponentAttributes := &AmazonManagedComponentAttributes{
//   	ComponentName: jsii.String("componentName"),
//
//   	// the properties below are optional
//   	ComponentVersion: jsii.String("componentVersion"),
//   }
//
// Experimental.
type AmazonManagedComponentAttributes struct {
	// The name of the Amazon-managed component.
	// Experimental.
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// The version of the Amazon-managed component.
	// Default: - the latest version of the component, x.x.x
	//
	// Experimental.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
}

