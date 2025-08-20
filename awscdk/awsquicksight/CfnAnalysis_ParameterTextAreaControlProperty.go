package awsquicksight


// A control to display a text box that is used to enter multiple entries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterTextAreaControlProperty := &ParameterTextAreaControlProperty{
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parametertextareacontrol.html
//
type CfnAnalysis_ParameterTextAreaControlProperty struct {
	// The ID of the `ParameterTextAreaControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parametertextareacontrol.html#cfn-quicksight-analysis-parametertextareacontrol-parametercontrolid
	//
	ParameterControlId *string `field:"required" json:"parameterControlId" yaml:"parameterControlId"`
	// The source parameter name of the `ParameterTextAreaControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parametertextareacontrol.html#cfn-quicksight-analysis-parametertextareacontrol-sourceparametername
	//
	SourceParameterName *string `field:"required" json:"sourceParameterName" yaml:"sourceParameterName"`
	// The title of the `ParameterTextAreaControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parametertextareacontrol.html#cfn-quicksight-analysis-parametertextareacontrol-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// The delimiter that is used to separate the lines in text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parametertextareacontrol.html#cfn-quicksight-analysis-parametertextareacontrol-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parametertextareacontrol.html#cfn-quicksight-analysis-parametertextareacontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

