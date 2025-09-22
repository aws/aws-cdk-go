package awsimagebuilder


// A reference to a ContainerRecipe resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerRecipeReference := &ContainerRecipeReference{
//   	ContainerRecipeArn: jsii.String("containerRecipeArn"),
//   }
//
type ContainerRecipeReference struct {
	// The Arn of the ContainerRecipe resource.
	ContainerRecipeArn *string `field:"required" json:"containerRecipeArn" yaml:"containerRecipeArn"`
}

