package previewawsquicksightmixins


// A control to display a horizontal toggle bar.
//
// This is used to change a value by sliding the toggle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterSliderControlProperty := &FilterSliderControlProperty{
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
//   	FilterControlId: jsii.String("filterControlId"),
//   	MaximumValue: jsii.Number(123),
//   	MinimumValue: jsii.Number(123),
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   	StepSize: jsii.Number(123),
//   	Title: jsii.String("title"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterslidercontrol.html
//
type CfnDashboardPropsMixin_FilterSliderControlProperty struct {
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterslidercontrol.html#cfn-quicksight-dashboard-filterslidercontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The ID of the `FilterSliderControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterslidercontrol.html#cfn-quicksight-dashboard-filterslidercontrol-filtercontrolid
	//
	FilterControlId *string `field:"optional" json:"filterControlId" yaml:"filterControlId"`
	// The larger value that is displayed at the right of the slider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterslidercontrol.html#cfn-quicksight-dashboard-filterslidercontrol-maximumvalue
	//
	// Default: - 0.
	//
	MaximumValue *float64 `field:"optional" json:"maximumValue" yaml:"maximumValue"`
	// The smaller value that is displayed at the left of the slider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterslidercontrol.html#cfn-quicksight-dashboard-filterslidercontrol-minimumvalue
	//
	// Default: - 0.
	//
	MinimumValue *float64 `field:"optional" json:"minimumValue" yaml:"minimumValue"`
	// The source filter ID of the `FilterSliderControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterslidercontrol.html#cfn-quicksight-dashboard-filterslidercontrol-sourcefilterid
	//
	SourceFilterId *string `field:"optional" json:"sourceFilterId" yaml:"sourceFilterId"`
	// The number of increments that the slider bar is divided into.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterslidercontrol.html#cfn-quicksight-dashboard-filterslidercontrol-stepsize
	//
	// Default: - 0.
	//
	StepSize *float64 `field:"optional" json:"stepSize" yaml:"stepSize"`
	// The title of the `FilterSliderControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterslidercontrol.html#cfn-quicksight-dashboard-filterslidercontrol-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The type of the `FilterSliderControl` . Choose one of the following options:.
	//
	// - `SINGLE_POINT` : Filter against(equals) a single data point.
	// - `RANGE` : Filter data that is in a specified range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterslidercontrol.html#cfn-quicksight-dashboard-filterslidercontrol-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

