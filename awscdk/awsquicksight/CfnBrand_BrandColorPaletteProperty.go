package awsquicksight


// The color palette.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   brandColorPaletteProperty := &BrandColorPaletteProperty{
//   	Accent: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   	Danger: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   	Dimension: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   	Info: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   	Measure: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   	Primary: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   	Secondary: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   	Success: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   	Warning: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html
//
type CfnBrand_BrandColorPaletteProperty struct {
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html#cfn-quicksight-brand-brandcolorpalette-accent
	//
	Accent interface{} `field:"optional" json:"accent" yaml:"accent"`
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html#cfn-quicksight-brand-brandcolorpalette-danger
	//
	Danger interface{} `field:"optional" json:"danger" yaml:"danger"`
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html#cfn-quicksight-brand-brandcolorpalette-dimension
	//
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html#cfn-quicksight-brand-brandcolorpalette-info
	//
	Info interface{} `field:"optional" json:"info" yaml:"info"`
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html#cfn-quicksight-brand-brandcolorpalette-measure
	//
	Measure interface{} `field:"optional" json:"measure" yaml:"measure"`
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html#cfn-quicksight-brand-brandcolorpalette-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html#cfn-quicksight-brand-brandcolorpalette-secondary
	//
	Secondary interface{} `field:"optional" json:"secondary" yaml:"secondary"`
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html#cfn-quicksight-brand-brandcolorpalette-success
	//
	Success interface{} `field:"optional" json:"success" yaml:"success"`
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandcolorpalette.html#cfn-quicksight-brand-brandcolorpalette-warning
	//
	Warning interface{} `field:"optional" json:"warning" yaml:"warning"`
}

