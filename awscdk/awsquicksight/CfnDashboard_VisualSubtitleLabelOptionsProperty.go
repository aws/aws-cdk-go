package awsquicksight


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
type CfnDashboard_VisualSubtitleLabelOptionsProperty struct {
	// `CfnDashboard.VisualSubtitleLabelOptionsProperty.FormatText`.
	FormatText interface{} `field:"optional" json:"formatText" yaml:"formatText"`
	// `CfnDashboard.VisualSubtitleLabelOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

