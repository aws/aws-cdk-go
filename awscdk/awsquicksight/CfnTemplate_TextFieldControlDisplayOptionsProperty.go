package awsquicksight


// The display options of a control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textFieldControlDisplayOptionsProperty := &TextFieldControlDisplayOptionsProperty{
//   	InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   		InfoIconText: jsii.String("infoIconText"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textfieldcontroldisplayoptions.html
//
type CfnTemplate_TextFieldControlDisplayOptionsProperty struct {
	// The configuration of info icon label options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textfieldcontroldisplayoptions.html#cfn-quicksight-template-textfieldcontroldisplayoptions-infoiconlabeloptions
	//
	InfoIconLabelOptions interface{} `field:"optional" json:"infoIconLabelOptions" yaml:"infoIconLabelOptions"`
	// The configuration of the placeholder options in a text field control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textfieldcontroldisplayoptions.html#cfn-quicksight-template-textfieldcontroldisplayoptions-placeholderoptions
	//
	PlaceholderOptions interface{} `field:"optional" json:"placeholderOptions" yaml:"placeholderOptions"`
	// The options to configure the title visibility, name, and font size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textfieldcontroldisplayoptions.html#cfn-quicksight-template-textfieldcontroldisplayoptions-titleoptions
	//
	TitleOptions interface{} `field:"optional" json:"titleOptions" yaml:"titleOptions"`
}

