package awsquicksight


// A control from a date parameter that specifies date and time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterDateTimePickerControlProperty := &ParameterDateTimePickerControlProperty{
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdatetimepickercontrol.html
//
type CfnTemplate_ParameterDateTimePickerControlProperty struct {
	// The ID of the `ParameterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdatetimepickercontrol.html#cfn-quicksight-template-parameterdatetimepickercontrol-parametercontrolid
	//
	ParameterControlId *string `field:"required" json:"parameterControlId" yaml:"parameterControlId"`
	// The name of the `ParameterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdatetimepickercontrol.html#cfn-quicksight-template-parameterdatetimepickercontrol-sourceparametername
	//
	SourceParameterName *string `field:"required" json:"sourceParameterName" yaml:"sourceParameterName"`
	// The title of the `ParameterDateTimePickerControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdatetimepickercontrol.html#cfn-quicksight-template-parameterdatetimepickercontrol-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdatetimepickercontrol.html#cfn-quicksight-template-parameterdatetimepickercontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

