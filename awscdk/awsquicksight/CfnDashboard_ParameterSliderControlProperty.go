package awsquicksight


// A control to display a horizontal toggle bar.
//
// This is used to change a value by sliding the toggle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterSliderControlProperty := &ParameterSliderControlProperty{
//   	MaximumValue: jsii.Number(123),
//   	MinimumValue: jsii.Number(123),
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   	StepSize: jsii.Number(123),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html
//
type CfnDashboard_ParameterSliderControlProperty struct {
	// The larger value that is displayed at the right of the slider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-maximumvalue
	//
	// Default: - 0.
	//
	MaximumValue *float64 `field:"required" json:"maximumValue" yaml:"maximumValue"`
	// The smaller value that is displayed at the left of the slider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-minimumvalue
	//
	// Default: - 0.
	//
	MinimumValue *float64 `field:"required" json:"minimumValue" yaml:"minimumValue"`
	// The ID of the `ParameterSliderControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-parametercontrolid
	//
	ParameterControlId *string `field:"required" json:"parameterControlId" yaml:"parameterControlId"`
	// The source parameter name of the `ParameterSliderControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-sourceparametername
	//
	SourceParameterName *string `field:"required" json:"sourceParameterName" yaml:"sourceParameterName"`
	// The number of increments that the slider bar is divided into.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-stepsize
	//
	// Default: - 0.
	//
	StepSize *float64 `field:"required" json:"stepSize" yaml:"stepSize"`
	// The title of the `ParameterSliderControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

