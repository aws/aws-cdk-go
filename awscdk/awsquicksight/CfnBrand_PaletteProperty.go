package awsquicksight


// The palette colors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paletteProperty := &PaletteProperty{
//   	Background: jsii.String("background"),
//   	Foreground: jsii.String("foreground"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-palette.html
//
type CfnBrand_PaletteProperty struct {
	// The background color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-palette.html#cfn-quicksight-brand-palette-background
	//
	Background *string `field:"optional" json:"background" yaml:"background"`
	// The foreground color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-palette.html#cfn-quicksight-brand-palette-foreground
	//
	Foreground *string `field:"optional" json:"foreground" yaml:"foreground"`
}

