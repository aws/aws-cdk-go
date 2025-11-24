package mixinsawsquicksight


// The display options of a control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dateIconVisibility interface{}
//   var helperTextVisibility interface{}
//
//   dateTimePickerControlDisplayOptionsProperty := &DateTimePickerControlDisplayOptionsProperty{
//   	DateIconVisibility: dateIconVisibility,
//   	DateTimeFormat: jsii.String("dateTimeFormat"),
//   	HelperTextVisibility: helperTextVisibility,
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimepickercontroldisplayoptions.html
//
type CfnTemplatePropsMixin_DateTimePickerControlDisplayOptionsProperty struct {
	// The date icon visibility of the `DateTimePickerControlDisplayOptions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimepickercontroldisplayoptions.html#cfn-quicksight-template-datetimepickercontroldisplayoptions-dateiconvisibility
	//
	DateIconVisibility interface{} `field:"optional" json:"dateIconVisibility" yaml:"dateIconVisibility"`
	// Customize how dates are formatted in controls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimepickercontroldisplayoptions.html#cfn-quicksight-template-datetimepickercontroldisplayoptions-datetimeformat
	//
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// The helper text visibility of the `DateTimePickerControlDisplayOptions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimepickercontroldisplayoptions.html#cfn-quicksight-template-datetimepickercontroldisplayoptions-helpertextvisibility
	//
	HelperTextVisibility interface{} `field:"optional" json:"helperTextVisibility" yaml:"helperTextVisibility"`
	// The configuration of info icon label options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimepickercontroldisplayoptions.html#cfn-quicksight-template-datetimepickercontroldisplayoptions-infoiconlabeloptions
	//
	InfoIconLabelOptions interface{} `field:"optional" json:"infoIconLabelOptions" yaml:"infoIconLabelOptions"`
	// The options to configure the title visibility, name, and font size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimepickercontroldisplayoptions.html#cfn-quicksight-template-datetimepickercontroldisplayoptions-titleoptions
	//
	TitleOptions interface{} `field:"optional" json:"titleOptions" yaml:"titleOptions"`
}

