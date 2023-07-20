package awspersonalize


// The AutoMLConfig object containing a list of recipes to search when AutoML is performed.
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

