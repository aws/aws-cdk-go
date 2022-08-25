package awsquicksight


// The theme configuration.
//
// This configuration contains all of the display properties for a theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   themeConfigurationProperty := &themeConfigurationProperty{
//   	dataColorPalette: &dataColorPaletteProperty{
//   		colors: []*string{
//   			jsii.String("colors"),
//   		},
//   		emptyFillColor: jsii.String("emptyFillColor"),
//   		minMaxGradient: []*string{
//   			jsii.String("minMaxGradient"),
//   		},
//   	},
//   	sheet: &sheetStyleProperty{
//   		tile: &tileStyleProperty{
//   			border: &borderStyleProperty{
//   				show: jsii.Boolean(false),
//   			},
//   		},
//   		tileLayout: &tileLayoutStyleProperty{
//   			gutter: &gutterStyleProperty{
//   				show: jsii.Boolean(false),
//   			},
//   			margin: &marginStyleProperty{
//   				show: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	typography: &typographyProperty{
//   		fontFamilies: []interface{}{
//   			&fontProperty{
//   				fontFamily: jsii.String("fontFamily"),
//   			},
//   		},
//   	},
//   	uiColorPalette: &uIColorPaletteProperty{
//   		accent: jsii.String("accent"),
//   		accentForeground: jsii.String("accentForeground"),
//   		danger: jsii.String("danger"),
//   		dangerForeground: jsii.String("dangerForeground"),
//   		dimension: jsii.String("dimension"),
//   		dimensionForeground: jsii.String("dimensionForeground"),
//   		measure: jsii.String("measure"),
//   		measureForeground: jsii.String("measureForeground"),
//   		primaryBackground: jsii.String("primaryBackground"),
//   		primaryForeground: jsii.String("primaryForeground"),
//   		secondaryBackground: jsii.String("secondaryBackground"),
//   		secondaryForeground: jsii.String("secondaryForeground"),
//   		success: jsii.String("success"),
//   		successForeground: jsii.String("successForeground"),
//   		warning: jsii.String("warning"),
//   		warningForeground: jsii.String("warningForeground"),
//   	},
//   }
//
type CfnTheme_ThemeConfigurationProperty struct {
	// Color properties that apply to chart data colors.
	DataColorPalette interface{} `field:"optional" json:"dataColorPalette" yaml:"dataColorPalette"`
	// Display options related to sheets.
	Sheet interface{} `field:"optional" json:"sheet" yaml:"sheet"`
	// `CfnTheme.ThemeConfigurationProperty.Typography`.
	Typography interface{} `field:"optional" json:"typography" yaml:"typography"`
	// Color properties that apply to the UI and to charts, excluding the colors that apply to data.
	UiColorPalette interface{} `field:"optional" json:"uiColorPalette" yaml:"uiColorPalette"`
}

