package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   typographyProperty := &typographyProperty{
//   	fontFamilies: []interface{}{
//   		&fontProperty{
//   			fontFamily: jsii.String("fontFamily"),
//   		},
//   	},
//   }
//
type CfnTheme_TypographyProperty struct {
	// `CfnTheme.TypographyProperty.FontFamilies`.
	FontFamilies interface{} `field:"optional" json:"fontFamilies" yaml:"fontFamilies"`
}

