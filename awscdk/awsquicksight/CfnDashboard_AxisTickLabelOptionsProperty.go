package awsquicksight


// The tick label options of an axis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisTickLabelOptionsProperty := &AxisTickLabelOptionsProperty{
//   	LabelOptions: &LabelOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   	RotationAngle: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisticklabeloptions.html
//
type CfnDashboard_AxisTickLabelOptionsProperty struct {
	// Determines whether or not the axis ticks are visible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisticklabeloptions.html#cfn-quicksight-dashboard-axisticklabeloptions-labeloptions
	//
	LabelOptions interface{} `field:"optional" json:"labelOptions" yaml:"labelOptions"`
	// The rotation angle of the axis tick labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisticklabeloptions.html#cfn-quicksight-dashboard-axisticklabeloptions-rotationangle
	//
	RotationAngle *float64 `field:"optional" json:"rotationAngle" yaml:"rotationAngle"`
}

