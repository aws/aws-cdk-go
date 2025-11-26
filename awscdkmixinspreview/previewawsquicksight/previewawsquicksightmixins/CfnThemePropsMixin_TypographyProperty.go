package previewawsquicksightmixins


// Determines the typography options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   typographyProperty := &TypographyProperty{
//   	FontFamilies: []interface{}{
//   		&FontProperty{
//   			FontFamily: jsii.String("fontFamily"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html
//
type CfnThemePropsMixin_TypographyProperty struct {
	// Determines the list of font families.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-fontfamilies
	//
	FontFamilies interface{} `field:"optional" json:"fontFamilies" yaml:"fontFamilies"`
}

