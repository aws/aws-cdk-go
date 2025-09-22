package awsdatabrew


// A reference to a Recipe resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recipeReference := &RecipeReference{
//   	RecipeName: jsii.String("recipeName"),
//   }
//
type RecipeReference struct {
	// The Name of the Recipe resource.
	RecipeName *string `field:"required" json:"recipeName" yaml:"recipeName"`
}

