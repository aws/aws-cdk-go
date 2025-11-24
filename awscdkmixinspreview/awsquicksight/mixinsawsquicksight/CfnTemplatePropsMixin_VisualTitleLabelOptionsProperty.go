package mixinsawsquicksight


// The title label options for a visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   visualTitleLabelOptionsProperty := &VisualTitleLabelOptionsProperty{
//   	FormatText: &ShortFormatTextProperty{
//   		PlainText: jsii.String("plainText"),
//   		RichText: jsii.String("richText"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualtitlelabeloptions.html
//
type CfnTemplatePropsMixin_VisualTitleLabelOptionsProperty struct {
	// The short text format of the title label, such as plain text or rich text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualtitlelabeloptions.html#cfn-quicksight-template-visualtitlelabeloptions-formattext
	//
	FormatText interface{} `field:"optional" json:"formatText" yaml:"formatText"`
	// The visibility of the title label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualtitlelabeloptions.html#cfn-quicksight-template-visualtitlelabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

