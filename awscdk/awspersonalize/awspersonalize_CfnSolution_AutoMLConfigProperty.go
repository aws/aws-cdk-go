package awspersonalize


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoMLConfigProperty := &autoMLConfigProperty{
//   	metricName: jsii.String("metricName"),
//   	recipeList: []*string{
//   		jsii.String("recipeList"),
//   	},
//   }
//
type CfnSolution_AutoMLConfigProperty struct {
	// `CfnSolution.AutoMLConfigProperty.MetricName`.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// `CfnSolution.AutoMLConfigProperty.RecipeList`.
	RecipeList *[]*string `field:"optional" json:"recipeList" yaml:"recipeList"`
}

