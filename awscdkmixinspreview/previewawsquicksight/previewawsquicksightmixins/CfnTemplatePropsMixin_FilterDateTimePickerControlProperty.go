package previewawsquicksightmixins


// A control from a date filter that is used to specify date and time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dateIconVisibility interface{}
//   var helperTextVisibility interface{}
//
//   filterDateTimePickerControlProperty := &FilterDateTimePickerControlProperty{
//   	CommitMode: jsii.String("commitMode"),
//   	DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   		DateIconVisibility: dateIconVisibility,
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
//   		HelperTextVisibility: helperTextVisibility,
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
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   	Title: jsii.String("title"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdatetimepickercontrol.html
//
type CfnTemplatePropsMixin_FilterDateTimePickerControlProperty struct {
	// The visibility configurationof the Apply button on a `DateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdatetimepickercontrol.html#cfn-quicksight-template-filterdatetimepickercontrol-commitmode
	//
	CommitMode *string `field:"optional" json:"commitMode" yaml:"commitMode"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdatetimepickercontrol.html#cfn-quicksight-template-filterdatetimepickercontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The ID of the `FilterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdatetimepickercontrol.html#cfn-quicksight-template-filterdatetimepickercontrol-filtercontrolid
	//
	FilterControlId *string `field:"optional" json:"filterControlId" yaml:"filterControlId"`
	// The source filter ID of the `FilterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdatetimepickercontrol.html#cfn-quicksight-template-filterdatetimepickercontrol-sourcefilterid
	//
	SourceFilterId *string `field:"optional" json:"sourceFilterId" yaml:"sourceFilterId"`
	// The title of the `FilterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdatetimepickercontrol.html#cfn-quicksight-template-filterdatetimepickercontrol-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The type of the `FilterDropDownControl` . Choose one of the following options:.
	//
	// - `MULTI_SELECT` : The user can select multiple entries from a dropdown menu.
	// - `SINGLE_SELECT` : The user can select a single entry from a dropdown menu.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdatetimepickercontrol.html#cfn-quicksight-template-filterdatetimepickercontrol-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

