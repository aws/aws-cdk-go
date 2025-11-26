package previewawsquicksightmixins


// A control from a date parameter that specifies date and time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parameterDateTimePickerControlProperty := &ParameterDateTimePickerControlProperty{
//   	DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   		DateIconVisibility: jsii.String("dateIconVisibility"),
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
//   		HelperTextVisibility: jsii.String("helperTextVisibility"),
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
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterdatetimepickercontrol.html
//
type CfnDashboardPropsMixin_ParameterDateTimePickerControlProperty struct {
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterdatetimepickercontrol.html#cfn-quicksight-dashboard-parameterdatetimepickercontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The ID of the `ParameterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterdatetimepickercontrol.html#cfn-quicksight-dashboard-parameterdatetimepickercontrol-parametercontrolid
	//
	ParameterControlId *string `field:"optional" json:"parameterControlId" yaml:"parameterControlId"`
	// The name of the `ParameterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterdatetimepickercontrol.html#cfn-quicksight-dashboard-parameterdatetimepickercontrol-sourceparametername
	//
	SourceParameterName *string `field:"optional" json:"sourceParameterName" yaml:"sourceParameterName"`
	// The title of the `ParameterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterdatetimepickercontrol.html#cfn-quicksight-dashboard-parameterdatetimepickercontrol-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

