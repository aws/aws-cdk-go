package interfacesawsimagebuilder


// A reference to a ImageRecipe resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageRecipeReference := &ImageRecipeReference{
//   	ImageRecipeArn: jsii.String("imageRecipeArn"),
//   }
//
type ImageRecipeReference struct {
	// The Arn of the ImageRecipe resource.
	ImageRecipeArn *string `field:"required" json:"imageRecipeArn" yaml:"imageRecipeArn"`
}

