package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTheme`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnThemeProps := &cfnThemeProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	themeId: jsii.String("themeId"),
//
//   	// the properties below are optional
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
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	versionDescription: jsii.String("versionDescription"),
//   }
//
type CfnThemeProps struct {
	// The ID of the AWS account where you want to store the new theme.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// An ID for the theme that you want to create.
	//
	// The theme ID is unique per AWS Region in each AWS account.
	ThemeId *string `field:"required" json:"themeId" yaml:"themeId"`
	// The ID of the theme that a custom theme will inherit from.
	//
	// All themes inherit from one of the starting themes defined by Amazon QuickSight. For a list of the starting themes, use `ListThemes` or choose *Themes* from within an analysis.
	BaseThemeId *string `field:"optional" json:"baseThemeId" yaml:"baseThemeId"`
	// The theme configuration, which contains the theme display properties.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A display name for the theme.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A valid grouping of resource permissions to apply to the new theme.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// A map of the key-value pairs for the resource tag or tags that you want to add to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A description of the first version of the theme that you're creating.
	//
	// Every time `UpdateTheme` is called, a new version is created. Each version of the theme has a description of the version in the `VersionDescription` field.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

