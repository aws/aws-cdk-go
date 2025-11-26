package previewawsquicksightmixins


// The sort configuration of a scatter plot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scatterPlotSortConfigurationProperty := &ScatterPlotSortConfigurationProperty{
//   	ScatterPlotLimitConfiguration: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-scatterplotsortconfiguration.html
//
type CfnAnalysisPropsMixin_ScatterPlotSortConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-scatterplotsortconfiguration.html#cfn-quicksight-analysis-scatterplotsortconfiguration-scatterplotlimitconfiguration
	//
	ScatterPlotLimitConfiguration interface{} `field:"optional" json:"scatterPlotLimitConfiguration" yaml:"scatterPlotLimitConfiguration"`
}

