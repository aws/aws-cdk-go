package awsquicksight


// The title label options for a visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualTitleLabelOptionsProperty := &VisualTitleLabelOptionsProperty{
//   	FormatText: &ShortFormatTextProperty{
//   		PlainText: jsii.String("plainText"),
//   		RichText: jsii.String("richText"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_VisualTitleLabelOptionsProperty struct {
	// The short text format of the title label, such as plain text or rich text.
	FormatText interface{} `field:"optional" json:"formatText" yaml:"formatText"`
	// The visibility of the title label.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

