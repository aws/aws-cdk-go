package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   themeVersionProperty := &themeVersionProperty{
//   	arn: jsii.String("arn"),
//   	baseThemeId: jsii.String("baseThemeId"),
//   	configuration: &themeConfigurationProperty{
//   		dataColorPalette: &dataColorPaletteProperty{
//   			colors: []*string{
//   				jsii.String("colors"),
//   			},
//   			emptyFillColor: jsii.String("emptyFillColor"),
//   			minMaxGradient: []*string{
//   				jsii.String("minMaxGradient"),
//   			},
//   		},
//   		sheet: &sheetStyleProperty{
//   			tile: &tileStyleProperty{
//   				border: &borderStyleProperty{
//   					show: jsii.Boolean(false),
//   				},
//   			},
//   			tileLayout: &tileLayoutStyleProperty{
//   				gutter: &gutterStyleProperty{
//   					show: jsii.Boolean(false),
//   				},
//   				margin: &marginStyleProperty{
//   					show: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		typography: &typographyProperty{
//   			fontFamilies: []interface{}{
//   				&fontProperty{
//   					fontFamily: jsii.String("fontFamily"),
//   				},
//   			},
//   		},
//   		uiColorPalette: &uIColorPaletteProperty{
//   			accent: jsii.String("accent"),
//   			accentForeground: jsii.String("accentForeground"),
//   			danger: jsii.String("danger"),
//   			dangerForeground: jsii.String("dangerForeground"),
//   			dimension: jsii.String("dimension"),
//   			dimensionForeground: jsii.String("dimensionForeground"),
//   			measure: jsii.String("measure"),
//   			measureForeground: jsii.String("measureForeground"),
//   			primaryBackground: jsii.String("primaryBackground"),
//   			primaryForeground: jsii.String("primaryForeground"),
//   			secondaryBackground: jsii.String("secondaryBackground"),
//   			secondaryForeground: jsii.String("secondaryForeground"),
//   			success: jsii.String("success"),
//   			successForeground: jsii.String("successForeground"),
//   			warning: jsii.String("warning"),
//   			warningForeground: jsii.String("warningForeground"),
//   		},
//   	},
//   	createdTime: jsii.String("createdTime"),
//   	description: jsii.String("description"),
//   	errors: []interface{}{
//   		&themeErrorProperty{
//   			message: jsii.String("message"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	status: jsii.String("status"),
//   	versionNumber: jsii.Number(123),
//   }
//
type CfnTheme_ThemeVersionProperty struct {
	// `CfnTheme.ThemeVersionProperty.Arn`.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// `CfnTheme.ThemeVersionProperty.BaseThemeId`.
	BaseThemeId *string `field:"optional" json:"baseThemeId" yaml:"baseThemeId"`
	// `CfnTheme.ThemeVersionProperty.Configuration`.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// `CfnTheme.ThemeVersionProperty.CreatedTime`.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// `CfnTheme.ThemeVersionProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnTheme.ThemeVersionProperty.Errors`.
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// `CfnTheme.ThemeVersionProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `CfnTheme.ThemeVersionProperty.VersionNumber`.
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

