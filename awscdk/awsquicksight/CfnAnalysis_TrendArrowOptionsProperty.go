package awsquicksight


// The options that determine the presentation of trend arrows in a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trendArrowOptionsProperty := &TrendArrowOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-trendarrowoptions.html
//
type CfnAnalysis_TrendArrowOptionsProperty struct {
	// The visibility of the trend arrows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-trendarrowoptions.html#cfn-quicksight-analysis-trendarrowoptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

