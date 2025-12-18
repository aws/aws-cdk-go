package previewawsconnectmixins


// Contains color palette configuration for different areas of a workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workspaceThemePaletteProperty := &WorkspaceThemePaletteProperty{
//   	Canvas: &PaletteCanvasProperty{
//   		ActiveBackground: jsii.String("activeBackground"),
//   		ContainerBackground: jsii.String("containerBackground"),
//   		PageBackground: jsii.String("pageBackground"),
//   	},
//   	Header: &PaletteHeaderProperty{
//   		Background: jsii.String("background"),
//   		InvertActionsColors: jsii.Boolean(false),
//   		Text: jsii.String("text"),
//   		TextHover: jsii.String("textHover"),
//   	},
//   	Navigation: &PaletteNavigationProperty{
//   		Background: jsii.String("background"),
//   		InvertActionsColors: jsii.Boolean(false),
//   		Text: jsii.String("text"),
//   		TextActive: jsii.String("textActive"),
//   		TextBackgroundActive: jsii.String("textBackgroundActive"),
//   		TextBackgroundHover: jsii.String("textBackgroundHover"),
//   		TextHover: jsii.String("textHover"),
//   	},
//   	Primary: &PalettePrimaryProperty{
//   		Active: jsii.String("active"),
//   		ContrastText: jsii.String("contrastText"),
//   		Default: jsii.String("default"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemepalette.html
//
type CfnWorkspacePropsMixin_WorkspaceThemePaletteProperty struct {
	// The color configuration for the canvas area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemepalette.html#cfn-connect-workspace-workspacethemepalette-canvas
	//
	Canvas interface{} `field:"optional" json:"canvas" yaml:"canvas"`
	// The color configuration for the header area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemepalette.html#cfn-connect-workspace-workspacethemepalette-header
	//
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// The color configuration for the navigation area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemepalette.html#cfn-connect-workspace-workspacethemepalette-navigation
	//
	Navigation interface{} `field:"optional" json:"navigation" yaml:"navigation"`
	// The primary color configuration used throughout the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemepalette.html#cfn-connect-workspace-workspacethemepalette-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

