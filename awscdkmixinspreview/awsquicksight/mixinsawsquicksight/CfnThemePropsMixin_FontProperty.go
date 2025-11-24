package mixinsawsquicksight


// Determines the font settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fontProperty := &FontProperty{
//   	FontFamily: jsii.String("fontFamily"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-font.html
//
type CfnThemePropsMixin_FontProperty struct {
	// Determines the font family settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-font.html#cfn-quicksight-theme-font-fontfamily
	//
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
}

