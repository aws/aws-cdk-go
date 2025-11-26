package previewawsquicksightmixins


// The default options that correspond to the `Slider` filter control type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultSliderControlOptionsProperty := &DefaultSliderControlOptionsProperty{
//   	DisplayOptions: &SliderControlDisplayOptionsProperty{
//   		InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   			InfoIconText: jsii.String("infoIconText"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TitleOptions: &LabelOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontFamily: jsii.String("fontFamily"),
//   				FontSize: &FontSizeProperty{
//   					Absolute: jsii.String("absolute"),
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	MaximumValue: jsii.Number(123),
//   	MinimumValue: jsii.Number(123),
//   	StepSize: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultslidercontroloptions.html
//
type CfnDashboardPropsMixin_DefaultSliderControlOptionsProperty struct {
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultslidercontroloptions.html#cfn-quicksight-dashboard-defaultslidercontroloptions-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The larger value that is displayed at the right of the slider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultslidercontroloptions.html#cfn-quicksight-dashboard-defaultslidercontroloptions-maximumvalue
	//
	// Default: - 0.
	//
	MaximumValue *float64 `field:"optional" json:"maximumValue" yaml:"maximumValue"`
	// The smaller value that is displayed at the left of the slider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultslidercontroloptions.html#cfn-quicksight-dashboard-defaultslidercontroloptions-minimumvalue
	//
	// Default: - 0.
	//
	MinimumValue *float64 `field:"optional" json:"minimumValue" yaml:"minimumValue"`
	// The number of increments that the slider bar is divided into.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultslidercontroloptions.html#cfn-quicksight-dashboard-defaultslidercontroloptions-stepsize
	//
	// Default: - 0.
	//
	StepSize *float64 `field:"optional" json:"stepSize" yaml:"stepSize"`
	// The type of the `DefaultSliderControlOptions` . Choose one of the following options:.
	//
	// - `SINGLE_POINT` : Filter against(equals) a single data point.
	// - `RANGE` : Filter data that is in a specified range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultslidercontroloptions.html#cfn-quicksight-dashboard-defaultslidercontroloptions-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

