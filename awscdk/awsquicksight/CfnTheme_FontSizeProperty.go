package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fontSizeProperty := &FontSizeProperty{
//   	Absolute: jsii.String("absolute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-fontsize.html
//
type CfnTheme_FontSizeProperty struct {
	// <p>The font size that you want to use in px.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-fontsize.html#cfn-quicksight-theme-fontsize-absolute
	//
	Absolute *string `field:"optional" json:"absolute" yaml:"absolute"`
}

