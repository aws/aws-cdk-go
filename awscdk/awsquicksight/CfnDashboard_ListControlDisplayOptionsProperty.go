package awsquicksight


// The display options of a control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listControlDisplayOptionsProperty := &ListControlDisplayOptionsProperty{
//   	InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   		InfoIconText: jsii.String("infoIconText"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	SearchOptions: &ListControlSearchOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	SelectAllOptions: &ListControlSelectAllOptionsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-listcontroldisplayoptions.html
//
type CfnDashboard_ListControlDisplayOptionsProperty struct {
	// The configuration of info icon label options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-listcontroldisplayoptions.html#cfn-quicksight-dashboard-listcontroldisplayoptions-infoiconlabeloptions
	//
	InfoIconLabelOptions interface{} `field:"optional" json:"infoIconLabelOptions" yaml:"infoIconLabelOptions"`
	// The configuration of the search options in a list control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-listcontroldisplayoptions.html#cfn-quicksight-dashboard-listcontroldisplayoptions-searchoptions
	//
	SearchOptions interface{} `field:"optional" json:"searchOptions" yaml:"searchOptions"`
	// The configuration of the `Select all` options in a list control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-listcontroldisplayoptions.html#cfn-quicksight-dashboard-listcontroldisplayoptions-selectalloptions
	//
	SelectAllOptions interface{} `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
	// The options to configure the title visibility, name, and font size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-listcontroldisplayoptions.html#cfn-quicksight-dashboard-listcontroldisplayoptions-titleoptions
	//
	TitleOptions interface{} `field:"optional" json:"titleOptions" yaml:"titleOptions"`
}

