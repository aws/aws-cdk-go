package previewawsquicksightmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnThemePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnThemeMixinProps := &CfnThemeMixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
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
//   			Tile: &TileStyleProperty{
//   				Border: &BorderStyleProperty{
//   					Show: jsii.Boolean(false),
//   				},
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
//   			FontFamilies: []interface{}{
//   				&FontProperty{
//   					FontFamily: jsii.String("fontFamily"),
//   				},
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
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   			Resource: jsii.String("resource"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThemeId: jsii.String("themeId"),
//   	VersionDescription: jsii.String("versionDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html
//
type CfnThemeMixinProps struct {
	// The ID of the AWS account where you want to store the new theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html#cfn-quicksight-theme-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The ID of the theme that a custom theme will inherit from.
	//
	// All themes inherit from one of the starting themes defined by Amazon Quick Sight. For a list of the starting themes, use `ListThemes` or choose *Themes* from within an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html#cfn-quicksight-theme-basethemeid
	//
	BaseThemeId *string `field:"optional" json:"baseThemeId" yaml:"baseThemeId"`
	// The theme configuration, which contains the theme display properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html#cfn-quicksight-theme-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A display name for the theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html#cfn-quicksight-theme-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A valid grouping of resource permissions to apply to the new theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html#cfn-quicksight-theme-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// A map of the key-value pairs for the resource tag or tags that you want to add to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html#cfn-quicksight-theme-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An ID for the theme that you want to create.
	//
	// The theme ID is unique per AWS Region in each AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html#cfn-quicksight-theme-themeid
	//
	ThemeId *string `field:"optional" json:"themeId" yaml:"themeId"`
	// A description of the first version of the theme that you're creating.
	//
	// Every time `UpdateTheme` is called, a new version is created. Each version of the theme has a description of the version in the `VersionDescription` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html#cfn-quicksight-theme-versiondescription
	//
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

