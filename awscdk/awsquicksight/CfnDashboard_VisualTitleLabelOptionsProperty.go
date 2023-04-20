package awsquicksight


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
	// `CfnDashboard.VisualTitleLabelOptionsProperty.FormatText`.
	FormatText interface{} `field:"optional" json:"formatText" yaml:"formatText"`
	// `CfnDashboard.VisualTitleLabelOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

