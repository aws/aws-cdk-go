package awsquicksight


// The options of a box plot visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   boxPlotOptionsProperty := &BoxPlotOptionsProperty{
//   	AllDataPointsVisibility: jsii.String("allDataPointsVisibility"),
//   	OutlierVisibility: jsii.String("outlierVisibility"),
//   	StyleOptions: &BoxPlotStyleOptionsProperty{
//   		FillStyle: jsii.String("fillStyle"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-boxplotoptions.html
//
type CfnDashboard_BoxPlotOptionsProperty struct {
	// Determines the visibility of all data points of the box plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-boxplotoptions.html#cfn-quicksight-dashboard-boxplotoptions-alldatapointsvisibility
	//
	AllDataPointsVisibility *string `field:"optional" json:"allDataPointsVisibility" yaml:"allDataPointsVisibility"`
	// Determines the visibility of the outlier in a box plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-boxplotoptions.html#cfn-quicksight-dashboard-boxplotoptions-outliervisibility
	//
	OutlierVisibility *string `field:"optional" json:"outlierVisibility" yaml:"outlierVisibility"`
	// The style options of the box plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-boxplotoptions.html#cfn-quicksight-dashboard-boxplotoptions-styleoptions
	//
	StyleOptions interface{} `field:"optional" json:"styleOptions" yaml:"styleOptions"`
}

