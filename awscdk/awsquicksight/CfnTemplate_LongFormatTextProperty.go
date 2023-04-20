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
type CfnTemplate_LongFormatTextProperty struct {
	// `CfnTemplate.LongFormatTextProperty.PlainText`.
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
	// `CfnTemplate.LongFormatTextProperty.RichText`.
	RichText *string `field:"optional" json:"richText" yaml:"richText"`
}

