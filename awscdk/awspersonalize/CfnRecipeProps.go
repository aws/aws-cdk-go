package awspersonalize


// Properties for defining a `CfnRecipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecipeProps := &CfnRecipeProps{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-recipe.html
//
type CfnRecipeProps struct {
	// The name of the recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-recipe.html#cfn-personalize-recipe-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

