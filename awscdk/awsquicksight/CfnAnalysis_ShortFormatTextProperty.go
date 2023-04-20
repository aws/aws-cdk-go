package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shortFormatTextProperty := &ShortFormatTextProperty{
//   	PlainText: jsii.String("plainText"),
//   	RichText: jsii.String("richText"),
//   }
//
type CfnAnalysis_ShortFormatTextProperty struct {
	// `CfnAnalysis.ShortFormatTextProperty.PlainText`.
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
	// `CfnAnalysis.ShortFormatTextProperty.RichText`.
	RichText *string `field:"optional" json:"richText" yaml:"richText"`
}

