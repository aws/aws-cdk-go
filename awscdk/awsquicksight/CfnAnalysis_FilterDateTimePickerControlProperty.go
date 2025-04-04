package awsquicksight


// A control from a date filter that is used to specify date and time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterDateTimePickerControlProperty := &FilterDateTimePickerControlProperty{
//   	FilterControlId: jsii.String("filterControlId"),
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	CommitMode: jsii.String("commitMode"),
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
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filterdatetimepickercontrol.html
//
type CfnAnalysis_FilterDateTimePickerControlProperty struct {
	// The ID of the `FilterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filterdatetimepickercontrol.html#cfn-quicksight-analysis-filterdatetimepickercontrol-filtercontrolid
	//
	FilterControlId *string `field:"required" json:"filterControlId" yaml:"filterControlId"`
	// The source filter ID of the `FilterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filterdatetimepickercontrol.html#cfn-quicksight-analysis-filterdatetimepickercontrol-sourcefilterid
	//
	SourceFilterId *string `field:"required" json:"sourceFilterId" yaml:"sourceFilterId"`
	// The title of the `FilterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filterdatetimepickercontrol.html#cfn-quicksight-analysis-filterdatetimepickercontrol-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// The visibility configurationof the Apply button on a `DateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filterdatetimepickercontrol.html#cfn-quicksight-analysis-filterdatetimepickercontrol-commitmode
	//
	CommitMode *string `field:"optional" json:"commitMode" yaml:"commitMode"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filterdatetimepickercontrol.html#cfn-quicksight-analysis-filterdatetimepickercontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The type of the `FilterDropDownControl` . Choose one of the following options:.
	//
	// - `MULTI_SELECT` : The user can select multiple entries from a dropdown menu.
	// - `SINGLE_SELECT` : The user can select a single entry from a dropdown menu.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filterdatetimepickercontrol.html#cfn-quicksight-analysis-filterdatetimepickercontrol-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

