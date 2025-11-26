package previewawsquicksightmixins


// The share label options for the labels.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   labelOptionsProperty := &LabelOptionsProperty{
//   	CustomLabel: jsii.String("customLabel"),
//   	FontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontFamily: jsii.String("fontFamily"),
//   		FontSize: &FontSizeProperty{
//   			Absolute: jsii.String("absolute"),
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-labeloptions.html
//
type CfnAnalysisPropsMixin_LabelOptionsProperty struct {
	// The text for the label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-labeloptions.html#cfn-quicksight-analysis-labeloptions-customlabel
	//
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The font configuration of the label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-labeloptions.html#cfn-quicksight-analysis-labeloptions-fontconfiguration
	//
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// Determines whether or not the label is visible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-labeloptions.html#cfn-quicksight-analysis-labeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

