package previewawsquicksightmixins


// The text format for the title.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   shortFormatTextProperty := &ShortFormatTextProperty{
//   	PlainText: jsii.String("plainText"),
//   	RichText: jsii.String("richText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-shortformattext.html
//
type CfnAnalysisPropsMixin_ShortFormatTextProperty struct {
	// Plain text format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-shortformattext.html#cfn-quicksight-analysis-shortformattext-plaintext
	//
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
	// Rich text.
	//
	// Examples of rich text include bold, underline, and italics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-shortformattext.html#cfn-quicksight-analysis-shortformattext-richtext
	//
	RichText *string `field:"optional" json:"richText" yaml:"richText"`
}

