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
type CfnTemplate_VisualSubtitleLabelOptionsProperty struct {
	// The long text format of the subtitle label, such as plain text or rich text.
	FormatText interface{} `field:"optional" json:"formatText" yaml:"formatText"`
	// The visibility of the subtitle label.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

