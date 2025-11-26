package previewawsquicksightmixins


// The options that determine the presentation of a waterfall visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   waterfallChartOptionsProperty := &WaterfallChartOptionsProperty{
//   	TotalBarLabel: jsii.String("totalBarLabel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartoptions.html
//
type CfnAnalysisPropsMixin_WaterfallChartOptionsProperty struct {
	// This option determines the total bar label of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartoptions.html#cfn-quicksight-analysis-waterfallchartoptions-totalbarlabel
	//
	TotalBarLabel *string `field:"optional" json:"totalBarLabel" yaml:"totalBarLabel"`
}

