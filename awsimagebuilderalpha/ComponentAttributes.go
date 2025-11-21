package awsimagebuilderalpha


// Properties for an EC2 Image Builder component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   componentAttributes := &ComponentAttributes{
//   	ComponentArn: jsii.String("componentArn"),
//   	ComponentName: jsii.String("componentName"),
//   	ComponentVersion: jsii.String("componentVersion"),
//   }
//
// Experimental.
type ComponentAttributes struct {
	// The ARN of the component.
	// Default: - the ARN is automatically constructed if a componentName is provided, otherwise a componentArn is
	// required.
	//
	// Experimental.
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// The name of the component.
	// Default: - the name is automatically constructed if a componentArn is provided, otherwise a componentName is
	// required.
	//
	// Experimental.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The version of the component.
	// Default: - the latest version of the component, x.x.x
	//
	// Experimental.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
}

