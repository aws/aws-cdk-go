package awspersonalize


// When the solution performs AutoML ( `performAutoML` is true in [CreateSolution](https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html) ), Amazon Personalize determines which recipe, from the specified list, optimizes the given metric. Amazon Personalize then uses that recipe for the solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoMLConfigProperty := &AutoMLConfigProperty{
//   	MetricName: jsii.String("metricName"),
//   	RecipeList: []*string{
//   		jsii.String("recipeList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-automlconfig.html
//
type CfnSolution_AutoMLConfigProperty struct {
	// The metric to optimize.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-automlconfig.html#cfn-personalize-solution-automlconfig-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The list of candidate recipes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-automlconfig.html#cfn-personalize-solution-automlconfig-recipelist
	//
	RecipeList *[]*string `field:"optional" json:"recipeList" yaml:"recipeList"`
}

