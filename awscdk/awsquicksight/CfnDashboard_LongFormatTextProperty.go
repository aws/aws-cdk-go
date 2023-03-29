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
type CfnDashboard_LongFormatTextProperty struct {
	// `CfnDashboard.LongFormatTextProperty.PlainText`.
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
	// `CfnDashboard.LongFormatTextProperty.RichText`.
	RichText *string `field:"optional" json:"richText" yaml:"richText"`
}

