package awsquicksight


// The subtitle label options for a visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualSubtitleLabelOptionsProperty := &VisualSubtitleLabelOptionsProperty{
//   	FormatText: &LongFormatTextProperty{
//   		PlainText: jsii.String("plainText"),
//   		RichText: jsii.String("richText"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualsubtitlelabeloptions.html
//
type CfnTemplate_VisualSubtitleLabelOptionsProperty struct {
	// The long text format of the subtitle label, such as plain text or rich text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualsubtitlelabeloptions.html#cfn-quicksight-template-visualsubtitlelabeloptions-formattext
	//
	FormatText interface{} `field:"optional" json:"formatText" yaml:"formatText"`
	// The visibility of the subtitle label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualsubtitlelabeloptions.html#cfn-quicksight-template-visualsubtitlelabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

