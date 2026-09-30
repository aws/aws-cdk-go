package awsquicksight


// The theme configuration.
//
// This configuration contains all of the display properties for a theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
//   		Background: &SheetBackgroundStyleProperty{
//   			Color: jsii.String("color"),
//   			Gradient: jsii.String("gradient"),
//   		},
//   		Tile: &TileStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &BorderStyleProperty{
//   				Color: jsii.String("color"),
//   				Show: jsii.Boolean(false),
//   				Width: jsii.String("width"),
//   			},
//   			BorderRadius: jsii.String("borderRadius"),
//   			Padding: jsii.String("padding"),
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
//   		AxisLabelFontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		AxisTitleFontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		DataLabelFontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		FontFamilies: []interface{}{
//   			&FontProperty{
//   				FontFamily: jsii.String("fontFamily"),
//   			},
//   		},
//   		LegendTitleFontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		LegendValueFontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		VisualSubtitleFontConfiguration: &VisualSubtitleFontConfigurationProperty{
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontFamily: jsii.String("fontFamily"),
//   				FontSize: &FontSizeProperty{
//   					Absolute: jsii.String("absolute"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			TextAlignment: jsii.String("textAlignment"),
//   			TextTransform: jsii.String("textTransform"),
//   		},
//   		VisualTitleFontConfiguration: &VisualTitleFontConfigurationProperty{
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontFamily: jsii.String("fontFamily"),
//   				FontSize: &FontSizeProperty{
//   					Absolute: jsii.String("absolute"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			TextAlignment: jsii.String("textAlignment"),
//   			TextTransform: jsii.String("textTransform"),
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

