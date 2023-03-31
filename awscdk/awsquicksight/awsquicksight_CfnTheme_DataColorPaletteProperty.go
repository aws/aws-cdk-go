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
//   dataColorPaletteProperty := &dataColorPaletteProperty{
//   	colors: []*string{
//   		jsii.String("colors"),
//   	},
//   	emptyFillColor: jsii.String("emptyFillColor"),
//   	minMaxGradient: []*string{
//   		jsii.String("minMaxGradient"),
//   	},
//   }
//
type CfnTheme_DataColorPaletteProperty struct {
	// The hexadecimal codes for the colors.
	Colors *[]*string `field:"optional" json:"colors" yaml:"colors"`
	// The hexadecimal code of a color that applies to charts where a lack of data is highlighted.
	EmptyFillColor *string `field:"optional" json:"emptyFillColor" yaml:"emptyFillColor"`
	// The minimum and maximum hexadecimal codes that describe a color gradient.
	MinMaxGradient *[]*string `field:"optional" json:"minMaxGradient" yaml:"minMaxGradient"`
}

