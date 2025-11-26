package previewawsquicksightmixins


// The display options for the axis label.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dataDriven interface{}
//
//   axisDisplayOptionsProperty := &AxisDisplayOptionsProperty{
//   	AxisLineVisibility: jsii.String("axisLineVisibility"),
//   	AxisOffset: jsii.String("axisOffset"),
//   	DataOptions: &AxisDataOptionsProperty{
//   		DateAxisOptions: &DateAxisOptionsProperty{
//   			MissingDateVisibility: jsii.String("missingDateVisibility"),
//   		},
//   		NumericAxisOptions: &NumericAxisOptionsProperty{
//   			Range: &AxisDisplayRangeProperty{
//   				DataDriven: dataDriven,
//   				MinMax: &AxisDisplayMinMaxRangeProperty{
//   					Maximum: jsii.Number(123),
//   					Minimum: jsii.Number(123),
//   				},
//   			},
//   			Scale: &AxisScaleProperty{
//   				Linear: &AxisLinearScaleProperty{
//   					StepCount: jsii.Number(123),
//   					StepSize: jsii.Number(123),
//   				},
//   				Logarithmic: &AxisLogarithmicScaleProperty{
//   					Base: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	GridLineVisibility: jsii.String("gridLineVisibility"),
//   	ScrollbarOptions: &ScrollBarOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   		VisibleRange: &VisibleRangeOptionsProperty{
//   			PercentRange: &PercentVisibleRangeProperty{
//   				From: jsii.Number(123),
//   				To: jsii.Number(123),
//   			},
//   		},
//   	},
//   	TickLabelOptions: &AxisTickLabelOptionsProperty{
//   		LabelOptions: &LabelOptionsProperty{
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
//   		RotationAngle: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-axisdisplayoptions.html
//
type CfnTemplatePropsMixin_AxisDisplayOptionsProperty struct {
	// Determines whether or not the axis line is visible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-axisdisplayoptions.html#cfn-quicksight-template-axisdisplayoptions-axislinevisibility
	//
	AxisLineVisibility *string `field:"optional" json:"axisLineVisibility" yaml:"axisLineVisibility"`
	// The offset value that determines the starting placement of the axis within a visual's bounds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-axisdisplayoptions.html#cfn-quicksight-template-axisdisplayoptions-axisoffset
	//
	AxisOffset *string `field:"optional" json:"axisOffset" yaml:"axisOffset"`
	// The data options for an axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-axisdisplayoptions.html#cfn-quicksight-template-axisdisplayoptions-dataoptions
	//
	DataOptions interface{} `field:"optional" json:"dataOptions" yaml:"dataOptions"`
	// Determines whether or not the grid line is visible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-axisdisplayoptions.html#cfn-quicksight-template-axisdisplayoptions-gridlinevisibility
	//
	GridLineVisibility *string `field:"optional" json:"gridLineVisibility" yaml:"gridLineVisibility"`
	// The scroll bar options for an axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-axisdisplayoptions.html#cfn-quicksight-template-axisdisplayoptions-scrollbaroptions
	//
	ScrollbarOptions interface{} `field:"optional" json:"scrollbarOptions" yaml:"scrollbarOptions"`
	// The tick label options of an axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-axisdisplayoptions.html#cfn-quicksight-template-axisdisplayoptions-ticklabeloptions
	//
	TickLabelOptions interface{} `field:"optional" json:"tickLabelOptions" yaml:"tickLabelOptions"`
}

