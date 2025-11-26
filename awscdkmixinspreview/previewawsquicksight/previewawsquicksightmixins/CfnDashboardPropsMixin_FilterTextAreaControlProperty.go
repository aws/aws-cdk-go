package previewawsquicksightmixins


// A control to display a text box that is used to enter multiple entries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterTextAreaControlProperty := &FilterTextAreaControlProperty{
//   	Delimiter: jsii.String("delimiter"),
//   	DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   		InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   			InfoIconText: jsii.String("infoIconText"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextareacontrol.html
//
type CfnDashboardPropsMixin_FilterTextAreaControlProperty struct {
	// The delimiter that is used to separate the lines in text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextareacontrol.html#cfn-quicksight-dashboard-filtertextareacontrol-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextareacontrol.html#cfn-quicksight-dashboard-filtertextareacontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The ID of the `FilterTextAreaControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextareacontrol.html#cfn-quicksight-dashboard-filtertextareacontrol-filtercontrolid
	//
	FilterControlId *string `field:"optional" json:"filterControlId" yaml:"filterControlId"`
	// The source filter ID of the `FilterTextAreaControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextareacontrol.html#cfn-quicksight-dashboard-filtertextareacontrol-sourcefilterid
	//
	SourceFilterId *string `field:"optional" json:"sourceFilterId" yaml:"sourceFilterId"`
	// The title of the `FilterTextAreaControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextareacontrol.html#cfn-quicksight-dashboard-filtertextareacontrol-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

