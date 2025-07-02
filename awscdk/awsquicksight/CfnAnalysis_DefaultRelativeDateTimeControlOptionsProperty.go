package awsquicksight


// The default options that correspond to the `RelativeDateTime` filter control type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultRelativeDateTimeControlOptionsProperty := &DefaultRelativeDateTimeControlOptionsProperty{
//   	CommitMode: jsii.String("commitMode"),
//   	DisplayOptions: &RelativeDateTimeControlDisplayOptionsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultrelativedatetimecontroloptions.html
//
type CfnAnalysis_DefaultRelativeDateTimeControlOptionsProperty struct {
	// The visibility configuration of the Apply button on a `RelativeDateTimeControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultrelativedatetimecontroloptions.html#cfn-quicksight-analysis-defaultrelativedatetimecontroloptions-commitmode
	//
	CommitMode *string `field:"optional" json:"commitMode" yaml:"commitMode"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultrelativedatetimecontroloptions.html#cfn-quicksight-analysis-defaultrelativedatetimecontroloptions-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

