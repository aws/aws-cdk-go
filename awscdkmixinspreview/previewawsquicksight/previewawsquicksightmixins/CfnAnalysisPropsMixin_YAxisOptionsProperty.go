package previewawsquicksightmixins


// The options that are available for a single Y axis in a chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   yAxisOptionsProperty := &YAxisOptionsProperty{
//   	YAxis: jsii.String("yAxis"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-yaxisoptions.html
//
type CfnAnalysisPropsMixin_YAxisOptionsProperty struct {
	// The Y axis type to be used in the chart.
	//
	// If you choose `PRIMARY_Y_AXIS` , the primary Y Axis is located on the leftmost vertical axis of the chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-yaxisoptions.html#cfn-quicksight-analysis-yaxisoptions-yaxis
	//
	YAxis *string `field:"optional" json:"yAxis" yaml:"yAxis"`
}

