package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   evaluationFormScoreThresholdProperty := &EvaluationFormScoreThresholdProperty{
//   	MaxScorePercentage: jsii.Number(123),
//   	MinScorePercentage: jsii.Number(123),
//   	PerformanceCategory: jsii.String("performanceCategory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformscorethreshold.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormScoreThresholdProperty struct {
	// The maximum score percentage for this threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformscorethreshold.html#cfn-connect-evaluationform-evaluationformscorethreshold-maxscorepercentage
	//
	MaxScorePercentage *float64 `field:"optional" json:"maxScorePercentage" yaml:"maxScorePercentage"`
	// The minimum score percentage for this threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformscorethreshold.html#cfn-connect-evaluationform-evaluationformscorethreshold-minscorepercentage
	//
	MinScorePercentage *float64 `field:"optional" json:"minScorePercentage" yaml:"minScorePercentage"`
	// The performance category name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformscorethreshold.html#cfn-connect-evaluationform-evaluationformscorethreshold-performancecategory
	//
	PerformanceCategory *string `field:"optional" json:"performanceCategory" yaml:"performanceCategory"`
}

