package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnTheme_TypographyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-fontfamilies
	//
	FontFamilies interface{} `field:"optional" json:"fontFamilies" yaml:"fontFamilies"`
}

