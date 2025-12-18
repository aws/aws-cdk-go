package previewawsconnectmixins


// Contains theme configuration for a workspace, supporting both light and dark modes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workspaceThemeProperty := &WorkspaceThemeProperty{
//   	Dark: &WorkspaceThemeConfigProperty{
//   		Palette: &WorkspaceThemePaletteProperty{
//   			Canvas: &PaletteCanvasProperty{
//   				ActiveBackground: jsii.String("activeBackground"),
//   				ContainerBackground: jsii.String("containerBackground"),
//   				PageBackground: jsii.String("pageBackground"),
//   			},
//   			Header: &PaletteHeaderProperty{
//   				Background: jsii.String("background"),
//   				InvertActionsColors: jsii.Boolean(false),
//   				Text: jsii.String("text"),
//   				TextHover: jsii.String("textHover"),
//   			},
//   			Navigation: &PaletteNavigationProperty{
//   				Background: jsii.String("background"),
//   				InvertActionsColors: jsii.Boolean(false),
//   				Text: jsii.String("text"),
//   				TextActive: jsii.String("textActive"),
//   				TextBackgroundActive: jsii.String("textBackgroundActive"),
//   				TextBackgroundHover: jsii.String("textBackgroundHover"),
//   				TextHover: jsii.String("textHover"),
//   			},
//   			Primary: &PalettePrimaryProperty{
//   				Active: jsii.String("active"),
//   				ContrastText: jsii.String("contrastText"),
//   				Default: jsii.String("default"),
//   			},
//   		},
//   		Typography: &WorkspaceThemeTypographyProperty{
//   			FontFamily: &FontFamilyProperty{
//   				Default: jsii.String("default"),
//   			},
//   		},
//   	},
//   	Light: &WorkspaceThemeConfigProperty{
//   		Palette: &WorkspaceThemePaletteProperty{
//   			Canvas: &PaletteCanvasProperty{
//   				ActiveBackground: jsii.String("activeBackground"),
//   				ContainerBackground: jsii.String("containerBackground"),
//   				PageBackground: jsii.String("pageBackground"),
//   			},
//   			Header: &PaletteHeaderProperty{
//   				Background: jsii.String("background"),
//   				InvertActionsColors: jsii.Boolean(false),
//   				Text: jsii.String("text"),
//   				TextHover: jsii.String("textHover"),
//   			},
//   			Navigation: &PaletteNavigationProperty{
//   				Background: jsii.String("background"),
//   				InvertActionsColors: jsii.Boolean(false),
//   				Text: jsii.String("text"),
//   				TextActive: jsii.String("textActive"),
//   				TextBackgroundActive: jsii.String("textBackgroundActive"),
//   				TextBackgroundHover: jsii.String("textBackgroundHover"),
//   				TextHover: jsii.String("textHover"),
//   			},
//   			Primary: &PalettePrimaryProperty{
//   				Active: jsii.String("active"),
//   				ContrastText: jsii.String("contrastText"),
//   				Default: jsii.String("default"),
//   			},
//   		},
//   		Typography: &WorkspaceThemeTypographyProperty{
//   			FontFamily: &FontFamilyProperty{
//   				Default: jsii.String("default"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacetheme.html
//
type CfnWorkspacePropsMixin_WorkspaceThemeProperty struct {
	// The theme configuration for dark mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacetheme.html#cfn-connect-workspace-workspacetheme-dark
	//
	Dark interface{} `field:"optional" json:"dark" yaml:"dark"`
	// The theme configuration for light mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacetheme.html#cfn-connect-workspace-workspacetheme-light
	//
	Light interface{} `field:"optional" json:"light" yaml:"light"`
}

