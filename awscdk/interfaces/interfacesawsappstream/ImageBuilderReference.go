package interfacesawsappstream


// A reference to a ImageBuilder resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageBuilderReference := &ImageBuilderReference{
//   	ImageBuilderName: jsii.String("imageBuilderName"),
//   }
//
type ImageBuilderReference struct {
	// The Name of the ImageBuilder resource.
	ImageBuilderName *string `field:"required" json:"imageBuilderName" yaml:"imageBuilderName"`
}

