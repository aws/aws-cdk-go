package mixinsawsquicksight


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotsortconfiguration.html
//
type CfnDashboardPropsMixin_ScatterPlotSortConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotsortconfiguration.html#cfn-quicksight-dashboard-scatterplotsortconfiguration-scatterplotlimitconfiguration
	//
	ScatterPlotLimitConfiguration interface{} `field:"optional" json:"scatterPlotLimitConfiguration" yaml:"scatterPlotLimitConfiguration"`
}

