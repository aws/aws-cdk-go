package awsquicksight


// A control to display a text box that is used to enter a single entry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterTextFieldControlProperty := &ParameterTextFieldControlProperty{
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	DisplayOptions: &TextFieldControlDisplayOptionsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametertextfieldcontrol.html
//
type CfnTemplate_ParameterTextFieldControlProperty struct {
	// The ID of the `ParameterTextFieldControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametertextfieldcontrol.html#cfn-quicksight-template-parametertextfieldcontrol-parametercontrolid
	//
	ParameterControlId *string `field:"required" json:"parameterControlId" yaml:"parameterControlId"`
	// The source parameter name of the `ParameterTextFieldControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametertextfieldcontrol.html#cfn-quicksight-template-parametertextfieldcontrol-sourceparametername
	//
	SourceParameterName *string `field:"required" json:"sourceParameterName" yaml:"sourceParameterName"`
	// The title of the `ParameterTextFieldControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametertextfieldcontrol.html#cfn-quicksight-template-parametertextfieldcontrol-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametertextfieldcontrol.html#cfn-quicksight-template-parametertextfieldcontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

