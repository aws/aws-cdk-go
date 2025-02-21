package awsquicksight


// The default options that correspond to the filter control type of a `DateTimePicker` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultDateTimePickerControlOptionsProperty := &DefaultDateTimePickerControlOptionsProperty{
//   	CommitMode: jsii.String("commitMode"),
//   	DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
//   		InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   			InfoIconText: jsii.String("infoIconText"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TitleOptions: &LabelOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
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
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-defaultdatetimepickercontroloptions.html
//
type CfnTemplate_DefaultDateTimePickerControlOptionsProperty struct {
	// The visibility configuration of the Apply button on a `DateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-defaultdatetimepickercontroloptions.html#cfn-quicksight-template-defaultdatetimepickercontroloptions-commitmode
	//
	CommitMode *string `field:"optional" json:"commitMode" yaml:"commitMode"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-defaultdatetimepickercontroloptions.html#cfn-quicksight-template-defaultdatetimepickercontroloptions-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The date time picker type of the `DefaultDateTimePickerControlOptions` . Choose one of the following options:.
	//
	// - `SINGLE_VALUED` : The filter condition is a fixed date.
	// - `DATE_RANGE` : The filter condition is a date time range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-defaultdatetimepickercontroloptions.html#cfn-quicksight-template-defaultdatetimepickercontroloptions-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

