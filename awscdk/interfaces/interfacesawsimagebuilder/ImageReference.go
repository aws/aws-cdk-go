package interfacesawsimagebuilder


// A reference to a Image resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageReference := &ImageReference{
//   	ImageArn: jsii.String("imageArn"),
//   }
//
type ImageReference struct {
	// The Arn of the Image resource.
	ImageArn *string `field:"required" json:"imageArn" yaml:"imageArn"`
}

