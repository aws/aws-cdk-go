package mixinsawsquicksight


// The theme configuration.
//
// This configuration contains all of the display properties for a theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   themeConfigurationProperty := &ThemeConfigurationProperty{
//   	DataColorPalette: &DataColorPaletteProperty{
//   		Colors: []*string{
//   			jsii.String("colors"),
//   		},
//   		EmptyFillColor: jsii.String("emptyFillColor"),
//   		MinMaxGradient: []*string{
//   			jsii.String("minMaxGradient"),
//   		},
//   	},
//   	Sheet: &SheetStyleProperty{
//   		Tile: &TileStyleProperty{
//   			Border: &BorderStyleProperty{
//   				Show: jsii.Boolean(false),
//   			},
//   		},
//   		TileLayout: &TileLayoutStyleProperty{
//   			Gutter: &GutterStyleProperty{
//   				Show: jsii.Boolean(false),
//   			},
//   			Margin: &MarginStyleProperty{
//   				Show: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Typography: &TypographyProperty{
//   		FontFamilies: []interface{}{
//   			&FontProperty{
//   				FontFamily: jsii.String("fontFamily"),
//   			},
//   		},
//   	},
//   	UiColorPalette: &UIColorPaletteProperty{
//   		Accent: jsii.String("accent"),
//   		AccentForeground: jsii.String("accentForeground"),
//   		Danger: jsii.String("danger"),
//   		DangerForeground: jsii.String("dangerForeground"),
//   		Dimension: jsii.String("dimension"),
//   		DimensionForeground: jsii.String("dimensionForeground"),
//   		Measure: jsii.String("measure"),
//   		MeasureForeground: jsii.String("measureForeground"),
//   		PrimaryBackground: jsii.String("primaryBackground"),
//   		PrimaryForeground: jsii.String("primaryForeground"),
//   		SecondaryBackground: jsii.String("secondaryBackground"),
//   		SecondaryForeground: jsii.String("secondaryForeground"),
//   		Success: jsii.String("success"),
//   		SuccessForeground: jsii.String("successForeground"),
//   		Warning: jsii.String("warning"),
//   		WarningForeground: jsii.String("warningForeground"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeconfiguration.html
//
type CfnThemePropsMixin_ThemeConfigurationProperty struct {
	// Color properties that apply to chart data colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeconfiguration.html#cfn-quicksight-theme-themeconfiguration-datacolorpalette
	//
	DataColorPalette interface{} `field:"optional" json:"dataColorPalette" yaml:"dataColorPalette"`
	// Display options related to sheets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeconfiguration.html#cfn-quicksight-theme-themeconfiguration-sheet
	//
	Sheet interface{} `field:"optional" json:"sheet" yaml:"sheet"`
	// Determines the typography options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeconfiguration.html#cfn-quicksight-theme-themeconfiguration-typography
	//
	Typography interface{} `field:"optional" json:"typography" yaml:"typography"`
	// Color properties that apply to the UI and to charts, excluding the colors that apply to data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeconfiguration.html#cfn-quicksight-theme-themeconfiguration-uicolorpalette
	//
	UiColorPalette interface{} `field:"optional" json:"uiColorPalette" yaml:"uiColorPalette"`
}

