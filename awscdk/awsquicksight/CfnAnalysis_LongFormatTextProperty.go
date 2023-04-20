package awsquicksight


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
type CfnAnalysis_LongFormatTextProperty struct {
	// `CfnAnalysis.LongFormatTextProperty.PlainText`.
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
	// `CfnAnalysis.LongFormatTextProperty.RichText`.
	RichText *string `field:"optional" json:"richText" yaml:"richText"`
}

