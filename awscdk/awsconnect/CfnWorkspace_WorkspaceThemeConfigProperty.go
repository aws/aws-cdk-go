package awsconnect


// Contains detailed theme configuration for a workspace, including colors, images, and typography.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspaceThemeConfigProperty := &WorkspaceThemeConfigProperty{
//   	Palette: &WorkspaceThemePaletteProperty{
//   		Canvas: &PaletteCanvasProperty{
//   			ActiveBackground: jsii.String("activeBackground"),
//   			ContainerBackground: jsii.String("containerBackground"),
//   			PageBackground: jsii.String("pageBackground"),
//   		},
//   		Header: &PaletteHeaderProperty{
//   			Background: jsii.String("background"),
//   			InvertActionsColors: jsii.Boolean(false),
//   			Text: jsii.String("text"),
//   			TextHover: jsii.String("textHover"),
//   		},
//   		Navigation: &PaletteNavigationProperty{
//   			Background: jsii.String("background"),
//   			InvertActionsColors: jsii.Boolean(false),
//   			Text: jsii.String("text"),
//   			TextActive: jsii.String("textActive"),
//   			TextBackgroundActive: jsii.String("textBackgroundActive"),
//   			TextBackgroundHover: jsii.String("textBackgroundHover"),
//   			TextHover: jsii.String("textHover"),
//   		},
//   		Primary: &PalettePrimaryProperty{
//   			Active: jsii.String("active"),
//   			ContrastText: jsii.String("contrastText"),
//   			Default: jsii.String("default"),
//   		},
//   	},
//   	Typography: &WorkspaceThemeTypographyProperty{
//   		FontFamily: &FontFamilyProperty{
//   			Default: jsii.String("default"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemeconfig.html
//
type CfnWorkspace_WorkspaceThemeConfigProperty struct {
	// The color palette configuration for the workspace theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemeconfig.html#cfn-connect-workspace-workspacethemeconfig-palette
	//
	Palette interface{} `field:"optional" json:"palette" yaml:"palette"`
	// The typography configuration for the workspace theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemeconfig.html#cfn-connect-workspace-workspacethemeconfig-typography
	//
	Typography interface{} `field:"optional" json:"typography" yaml:"typography"`
}

