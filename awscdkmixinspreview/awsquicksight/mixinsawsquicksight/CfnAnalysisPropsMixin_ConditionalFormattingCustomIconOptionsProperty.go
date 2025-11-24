package mixinsawsquicksight


// Custom icon options for an icon set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionalFormattingCustomIconOptionsProperty := &ConditionalFormattingCustomIconOptionsProperty{
//   	Icon: jsii.String("icon"),
//   	UnicodeIcon: jsii.String("unicodeIcon"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcustomiconoptions.html
//
type CfnAnalysisPropsMixin_ConditionalFormattingCustomIconOptionsProperty struct {
	// Determines the type of icon.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcustomiconoptions.html#cfn-quicksight-analysis-conditionalformattingcustomiconoptions-icon
	//
	Icon *string `field:"optional" json:"icon" yaml:"icon"`
	// Determines the Unicode icon type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingcustomiconoptions.html#cfn-quicksight-analysis-conditionalformattingcustomiconoptions-unicodeicon
	//
	UnicodeIcon *string `field:"optional" json:"unicodeIcon" yaml:"unicodeIcon"`
}

