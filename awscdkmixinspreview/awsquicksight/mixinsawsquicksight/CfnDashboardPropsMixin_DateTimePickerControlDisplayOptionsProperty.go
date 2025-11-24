package mixinsawsquicksight


// The display options of a control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dateTimePickerControlDisplayOptionsProperty := &DateTimePickerControlDisplayOptionsProperty{
//   	DateIconVisibility: jsii.String("dateIconVisibility"),
//   	DateTimeFormat: jsii.String("dateTimeFormat"),
//   	HelperTextVisibility: jsii.String("helperTextVisibility"),
//   	InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   		InfoIconText: jsii.String("infoIconText"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	TitleOptions: &LabelOptionsProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html
//
type CfnDashboardPropsMixin_DateTimePickerControlDisplayOptionsProperty struct {
	// The date icon visibility of the `DateTimePickerControlDisplayOptions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html#cfn-quicksight-dashboard-datetimepickercontroldisplayoptions-dateiconvisibility
	//
	DateIconVisibility *string `field:"optional" json:"dateIconVisibility" yaml:"dateIconVisibility"`
	// Customize how dates are formatted in controls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html#cfn-quicksight-dashboard-datetimepickercontroldisplayoptions-datetimeformat
	//
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// The helper text visibility of the `DateTimePickerControlDisplayOptions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html#cfn-quicksight-dashboard-datetimepickercontroldisplayoptions-helpertextvisibility
	//
	HelperTextVisibility *string `field:"optional" json:"helperTextVisibility" yaml:"helperTextVisibility"`
	// The configuration of info icon label options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html#cfn-quicksight-dashboard-datetimepickercontroldisplayoptions-infoiconlabeloptions
	//
	InfoIconLabelOptions interface{} `field:"optional" json:"infoIconLabelOptions" yaml:"infoIconLabelOptions"`
	// The options to configure the title visibility, name, and font size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html#cfn-quicksight-dashboard-datetimepickercontroldisplayoptions-titleoptions
	//
	TitleOptions interface{} `field:"optional" json:"titleOptions" yaml:"titleOptions"`
}

