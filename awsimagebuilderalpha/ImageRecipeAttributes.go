package awsimagebuilderalpha


// Properties for an EC2 Image Builder image recipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   imageRecipeAttributes := &ImageRecipeAttributes{
//   	ImageRecipeArn: jsii.String("imageRecipeArn"),
//   	ImageRecipeName: jsii.String("imageRecipeName"),
//   	ImageRecipeVersion: jsii.String("imageRecipeVersion"),
//   }
//
// Experimental.
type ImageRecipeAttributes struct {
	// The ARN of the image recipe.
	// Default: - derived from the imageRecipeName.
	//
	// Experimental.
	ImageRecipeArn *string `field:"optional" json:"imageRecipeArn" yaml:"imageRecipeArn"`
	// The name of the image recipe.
	// Default: - derived from the imageRecipeArn.
	//
	// Experimental.
	ImageRecipeName *string `field:"optional" json:"imageRecipeName" yaml:"imageRecipeName"`
	// The version of the image recipe.
	// Default: - derived from imageRecipeArn. if a imageRecipeName is provided, the latest version, x.x.x, will
	// be used.
	//
	// Experimental.
	ImageRecipeVersion *string `field:"optional" json:"imageRecipeVersion" yaml:"imageRecipeVersion"`
}

