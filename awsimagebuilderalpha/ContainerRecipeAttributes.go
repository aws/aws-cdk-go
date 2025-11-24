package awsimagebuilderalpha


// Properties for an EC2 Image Builder container recipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   containerRecipeAttributes := &ContainerRecipeAttributes{
//   	ContainerRecipeArn: jsii.String("containerRecipeArn"),
//   	ContainerRecipeName: jsii.String("containerRecipeName"),
//   	ContainerRecipeVersion: jsii.String("containerRecipeVersion"),
//   }
//
// Experimental.
type ContainerRecipeAttributes struct {
	// The ARN of the container recipe.
	// Default: - derived from containerRecipeName.
	//
	// Experimental.
	ContainerRecipeArn *string `field:"optional" json:"containerRecipeArn" yaml:"containerRecipeArn"`
	// The name of the container recipe.
	// Default: - derived from containerRecipeArn.
	//
	// Experimental.
	ContainerRecipeName *string `field:"optional" json:"containerRecipeName" yaml:"containerRecipeName"`
	// The version of the container recipe.
	// Default: - derived from containerRecipeArn. if a containerRecipeName is provided, the latest version, x.x.x, will
	// be used.
	//
	// Experimental.
	ContainerRecipeVersion *string `field:"optional" json:"containerRecipeVersion" yaml:"containerRecipeVersion"`
}

