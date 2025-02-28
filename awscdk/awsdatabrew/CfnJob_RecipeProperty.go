package awsdatabrew


// Represents one or more actions to be performed on a DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recipeProperty := &RecipeProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-recipe.html
//
type CfnJob_RecipeProperty struct {
	// The unique name for the recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-recipe.html#cfn-databrew-job-recipe-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier for the version for the recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-recipe.html#cfn-databrew-job-recipe-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

