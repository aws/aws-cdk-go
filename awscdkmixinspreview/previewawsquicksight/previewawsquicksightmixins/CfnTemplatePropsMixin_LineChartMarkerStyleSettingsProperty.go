package previewawsquicksightmixins


// Marker styles options for a line series in `LineChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lineChartMarkerStyleSettingsProperty := &LineChartMarkerStyleSettingsProperty{
//   	MarkerColor: jsii.String("markerColor"),
//   	MarkerShape: jsii.String("markerShape"),
//   	MarkerSize: jsii.String("markerSize"),
//   	MarkerVisibility: jsii.String("markerVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html
//
type CfnTemplatePropsMixin_LineChartMarkerStyleSettingsProperty struct {
	// Color of marker in the series.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html#cfn-quicksight-template-linechartmarkerstylesettings-markercolor
	//
	MarkerColor *string `field:"optional" json:"markerColor" yaml:"markerColor"`
	// Shape option for markers in the series.
	//
	// - `CIRCLE` : Show marker as a circle.
	// - `TRIANGLE` : Show marker as a triangle.
	// - `SQUARE` : Show marker as a square.
	// - `DIAMOND` : Show marker as a diamond.
	// - `ROUNDED_SQUARE` : Show marker as a rounded square.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html#cfn-quicksight-template-linechartmarkerstylesettings-markershape
	//
	MarkerShape *string `field:"optional" json:"markerShape" yaml:"markerShape"`
	// Size of marker in the series.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html#cfn-quicksight-template-linechartmarkerstylesettings-markersize
	//
	MarkerSize *string `field:"optional" json:"markerSize" yaml:"markerSize"`
	// Configuration option that determines whether to show the markers in the series.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html#cfn-quicksight-template-linechartmarkerstylesettings-markervisibility
	//
	MarkerVisibility *string `field:"optional" json:"markerVisibility" yaml:"markerVisibility"`
}

