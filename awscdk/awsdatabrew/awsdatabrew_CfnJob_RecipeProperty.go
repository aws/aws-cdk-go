package awsdatabrew


// Represents one or more actions to be performed on a DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recipeProperty := &recipeProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	version: jsii.String("version"),
//   }
//
type CfnJob_RecipeProperty struct {
	// The unique name for the recipe.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier for the version for the recipe.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

