package awsquicksight


// The theme colors that are used for data colors in charts.
//
// The colors description is a hexadecimal color code that consists of six alphanumerical characters, prefixed with `#` , for example #37BFF5.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataColorPaletteProperty := &DataColorPaletteProperty{
//   	Colors: []*string{
//   		jsii.String("colors"),
//   	},
//   	EmptyFillColor: jsii.String("emptyFillColor"),
//   	MinMaxGradient: []*string{
//   		jsii.String("minMaxGradient"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html
//
type CfnTheme_DataColorPaletteProperty struct {
	// The hexadecimal codes for the colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html#cfn-quicksight-theme-datacolorpalette-colors
	//
	Colors *[]*string `field:"optional" json:"colors" yaml:"colors"`
	// The hexadecimal code of a color that applies to charts where a lack of data is highlighted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html#cfn-quicksight-theme-datacolorpalette-emptyfillcolor
	//
	EmptyFillColor *string `field:"optional" json:"emptyFillColor" yaml:"emptyFillColor"`
	// The minimum and maximum hexadecimal codes that describe a color gradient.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html#cfn-quicksight-theme-datacolorpalette-minmaxgradient
	//
	MinMaxGradient *[]*string `field:"optional" json:"minMaxGradient" yaml:"minMaxGradient"`
}

