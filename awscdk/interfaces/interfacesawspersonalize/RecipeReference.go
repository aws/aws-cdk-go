package interfacesawspersonalize


// A reference to a Recipe resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recipeReference := &RecipeReference{
//   	RecipeArn: jsii.String("recipeArn"),
//   }
//
type RecipeReference struct {
	// The RecipeArn of the Recipe resource.
	RecipeArn *string `field:"required" json:"recipeArn" yaml:"recipeArn"`
}

