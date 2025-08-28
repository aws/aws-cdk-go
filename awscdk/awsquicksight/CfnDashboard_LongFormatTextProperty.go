package awsquicksight


// The text format for a subtitle.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   longFormatTextProperty := &LongFormatTextProperty{
//   	PlainText: jsii.String("plainText"),
//   	RichText: jsii.String("richText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-longformattext.html
//
type CfnDashboard_LongFormatTextProperty struct {
	// Plain text format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-longformattext.html#cfn-quicksight-dashboard-longformattext-plaintext
	//
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
	// Rich text.
	//
	// Examples of rich text include bold, underline, and italics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-longformattext.html#cfn-quicksight-dashboard-longformattext-richtext
	//
	RichText *string `field:"optional" json:"richText" yaml:"richText"`
}

