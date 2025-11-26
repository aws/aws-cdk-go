package previewawsquicksightmixins


// The options that determine the presentation of the data labels.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   funnelChartDataLabelOptionsProperty := &FunnelChartDataLabelOptionsProperty{
//   	CategoryLabelVisibility: jsii.String("categoryLabelVisibility"),
//   	LabelColor: jsii.String("labelColor"),
//   	LabelFontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontFamily: jsii.String("fontFamily"),
//   		FontSize: &FontSizeProperty{
//   			Absolute: jsii.String("absolute"),
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	MeasureDataLabelStyle: jsii.String("measureDataLabelStyle"),
//   	MeasureLabelVisibility: jsii.String("measureLabelVisibility"),
//   	Position: jsii.String("position"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartdatalabeloptions.html
//
type CfnAnalysisPropsMixin_FunnelChartDataLabelOptionsProperty struct {
	// The visibility of the category labels within the data labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartdatalabeloptions.html#cfn-quicksight-analysis-funnelchartdatalabeloptions-categorylabelvisibility
	//
	CategoryLabelVisibility *string `field:"optional" json:"categoryLabelVisibility" yaml:"categoryLabelVisibility"`
	// The color of the data label text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartdatalabeloptions.html#cfn-quicksight-analysis-funnelchartdatalabeloptions-labelcolor
	//
	LabelColor *string `field:"optional" json:"labelColor" yaml:"labelColor"`
	// The font configuration for the data labels.
	//
	// Only the `FontSize` attribute of the font configuration is used for data labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartdatalabeloptions.html#cfn-quicksight-analysis-funnelchartdatalabeloptions-labelfontconfiguration
	//
	LabelFontConfiguration interface{} `field:"optional" json:"labelFontConfiguration" yaml:"labelFontConfiguration"`
	// Determines the style of the metric labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartdatalabeloptions.html#cfn-quicksight-analysis-funnelchartdatalabeloptions-measuredatalabelstyle
	//
	MeasureDataLabelStyle *string `field:"optional" json:"measureDataLabelStyle" yaml:"measureDataLabelStyle"`
	// The visibility of the measure labels within the data labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartdatalabeloptions.html#cfn-quicksight-analysis-funnelchartdatalabeloptions-measurelabelvisibility
	//
	MeasureLabelVisibility *string `field:"optional" json:"measureLabelVisibility" yaml:"measureLabelVisibility"`
	// Determines the positioning of the data label relative to a section of the funnel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartdatalabeloptions.html#cfn-quicksight-analysis-funnelchartdatalabeloptions-position
	//
	Position *string `field:"optional" json:"position" yaml:"position"`
	// The visibility option that determines if data labels are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartdatalabeloptions.html#cfn-quicksight-analysis-funnelchartdatalabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

