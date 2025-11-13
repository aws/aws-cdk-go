package interfacesawssagemaker


// A reference to a ImageVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageVersionReference := &ImageVersionReference{
//   	ImageVersionArn: jsii.String("imageVersionArn"),
//   }
//
type ImageVersionReference struct {
	// The ImageVersionArn of the ImageVersion resource.
	ImageVersionArn *string `field:"required" json:"imageVersionArn" yaml:"imageVersionArn"`
}

