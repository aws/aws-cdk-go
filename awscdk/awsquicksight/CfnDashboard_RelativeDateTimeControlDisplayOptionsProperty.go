package awsquicksight


// The display options of a control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relativeDateTimeControlDisplayOptionsProperty := &RelativeDateTimeControlDisplayOptionsProperty{
//   	DateTimeFormat: jsii.String("dateTimeFormat"),
//   	TitleOptions: &LabelOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontSize: &FontSizeProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatetimecontroldisplayoptions.html
//
type CfnDashboard_RelativeDateTimeControlDisplayOptionsProperty struct {
	// Customize how dates are formatted in controls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatetimecontroldisplayoptions.html#cfn-quicksight-dashboard-relativedatetimecontroldisplayoptions-datetimeformat
	//
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// The options to configure the title visibility, name, and font size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatetimecontroldisplayoptions.html#cfn-quicksight-dashboard-relativedatetimecontroldisplayoptions-titleoptions
	//
	TitleOptions interface{} `field:"optional" json:"titleOptions" yaml:"titleOptions"`
}

