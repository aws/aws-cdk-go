package awsquicksight


// A version of a theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   themeVersionProperty := &ThemeVersionProperty{
//   	Arn: jsii.String("arn"),
//   	BaseThemeId: jsii.String("baseThemeId"),
//   	Configuration: &ThemeConfigurationProperty{
//   		DataColorPalette: &DataColorPaletteProperty{
//   			Colors: []*string{
//   				jsii.String("colors"),
//   			},
//   			EmptyFillColor: jsii.String("emptyFillColor"),
//   			MinMaxGradient: []*string{
//   				jsii.String("minMaxGradient"),
//   			},
//   		},
//   		Sheet: &SheetStyleProperty{
//   			Background: &SheetBackgroundStyleProperty{
//   				Color: jsii.String("color"),
//   				Gradient: jsii.String("gradient"),
//   			},
//   			Tile: &TileStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &BorderStyleProperty{
//   					Color: jsii.String("color"),
//   					Show: jsii.Boolean(false),
//   					Width: jsii.String("width"),
//   				},
//   				BorderRadius: jsii.String("borderRadius"),
//   				Padding: jsii.String("padding"),
//   			},
//   			TileLayout: &TileLayoutStyleProperty{
//   				Gutter: &GutterStyleProperty{
//   					Show: jsii.Boolean(false),
//   				},
//   				Margin: &MarginStyleProperty{
//   					Show: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		Typography: &TypographyProperty{
//   			AxisLabelFontConfiguration: &FontConfigurationProperty{
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
//   			AxisTitleFontConfiguration: &FontConfigurationProperty{
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
//   			DataLabelFontConfiguration: &FontConfigurationProperty{
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
//   			FontFamilies: []interface{}{
//   				&FontProperty{
//   					FontFamily: jsii.String("fontFamily"),
//   				},
//   			},
//   			LegendTitleFontConfiguration: &FontConfigurationProperty{
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
//   			LegendValueFontConfiguration: &FontConfigurationProperty{
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
//   			VisualSubtitleFontConfiguration: &VisualSubtitleFontConfigurationProperty{
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				TextAlignment: jsii.String("textAlignment"),
//   				TextTransform: jsii.String("textTransform"),
//   			},
//   			VisualTitleFontConfiguration: &VisualTitleFontConfigurationProperty{
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				TextAlignment: jsii.String("textAlignment"),
//   				TextTransform: jsii.String("textTransform"),
//   			},
//   		},
//   		UiColorPalette: &UIColorPaletteProperty{
//   			Accent: jsii.String("accent"),
//   			AccentForeground: jsii.String("accentForeground"),
//   			Danger: jsii.String("danger"),
//   			DangerForeground: jsii.String("dangerForeground"),
//   			Dimension: jsii.String("dimension"),
//   			DimensionForeground: jsii.String("dimensionForeground"),
//   			Measure: jsii.String("measure"),
//   			MeasureForeground: jsii.String("measureForeground"),
//   			PrimaryBackground: jsii.String("primaryBackground"),
//   			PrimaryForeground: jsii.String("primaryForeground"),
//   			SecondaryBackground: jsii.String("secondaryBackground"),
//   			SecondaryForeground: jsii.String("secondaryForeground"),
//   			Success: jsii.String("success"),
//   			SuccessForeground: jsii.String("successForeground"),
//   			Warning: jsii.String("warning"),
//   			WarningForeground: jsii.String("warningForeground"),
//   		},
//   	},
//   	CreatedTime: jsii.String("createdTime"),
//   	Description: jsii.String("description"),
//   	Errors: []interface{}{
//   		&ThemeErrorProperty{
//   			Message: jsii.String("message"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	VersionNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeversion.html
//
type CfnThemePropsMixin_ThemeVersionProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeversion.html#cfn-quicksight-theme-themeversion-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The Quick Sight-defined ID of the theme that a custom theme inherits from.
	//
	// All themes initially inherit from a default Quick Sight theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeversion.html#cfn-quicksight-theme-themeversion-basethemeid
	//
	BaseThemeId *string `field:"optional" json:"baseThemeId" yaml:"baseThemeId"`
	// The theme configuration, which contains all the theme display properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeversion.html#cfn-quicksight-theme-themeversion-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The date and time that this theme version was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeversion.html#cfn-quicksight-theme-themeversion-createdtime
	//
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// The description of the theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeversion.html#cfn-quicksight-theme-themeversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Errors associated with the theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeversion.html#cfn-quicksight-theme-themeversion-errors
	//
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// The status of the theme version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeversion.html#cfn-quicksight-theme-themeversion-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The version number of the theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeversion.html#cfn-quicksight-theme-themeversion-versionnumber
	//
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

