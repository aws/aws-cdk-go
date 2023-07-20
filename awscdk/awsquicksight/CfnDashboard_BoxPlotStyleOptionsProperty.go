package awsquicksight


// The style options of the box plot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   boxPlotStyleOptionsProperty := &BoxPlotStyleOptionsProperty{
//   	FillStyle: jsii.String("fillStyle"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-boxplotstyleoptions.html
//
type CfnDashboard_BoxPlotStyleOptionsProperty struct {
	// The fill styles (solid, transparent) of the box plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-boxplotstyleoptions.html#cfn-quicksight-dashboard-boxplotstyleoptions-fillstyle
	//
	FillStyle *string `field:"optional" json:"fillStyle" yaml:"fillStyle"`
}

