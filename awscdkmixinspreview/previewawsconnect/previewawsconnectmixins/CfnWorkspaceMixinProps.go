package previewawsconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWorkspacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkspaceMixinProps := &CfnWorkspaceMixinProps{
//   	Associations: []*string{
//   		jsii.String("associations"),
//   	},
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Media: []interface{}{
//   		&MediaItemProperty{
//   			Source: jsii.String("source"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Pages: []interface{}{
//   		&WorkspacePageProperty{
//   			InputData: jsii.String("inputData"),
//   			Page: jsii.String("page"),
//   			ResourceArn: jsii.String("resourceArn"),
//   			Slug: jsii.String("slug"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Theme: &WorkspaceThemeProperty{
//   		Dark: &WorkspaceThemeConfigProperty{
//   			Palette: &WorkspaceThemePaletteProperty{
//   				Canvas: &PaletteCanvasProperty{
//   					ActiveBackground: jsii.String("activeBackground"),
//   					ContainerBackground: jsii.String("containerBackground"),
//   					PageBackground: jsii.String("pageBackground"),
//   				},
//   				Header: &PaletteHeaderProperty{
//   					Background: jsii.String("background"),
//   					InvertActionsColors: jsii.Boolean(false),
//   					Text: jsii.String("text"),
//   					TextHover: jsii.String("textHover"),
//   				},
//   				Navigation: &PaletteNavigationProperty{
//   					Background: jsii.String("background"),
//   					InvertActionsColors: jsii.Boolean(false),
//   					Text: jsii.String("text"),
//   					TextActive: jsii.String("textActive"),
//   					TextBackgroundActive: jsii.String("textBackgroundActive"),
//   					TextBackgroundHover: jsii.String("textBackgroundHover"),
//   					TextHover: jsii.String("textHover"),
//   				},
//   				Primary: &PalettePrimaryProperty{
//   					Active: jsii.String("active"),
//   					ContrastText: jsii.String("contrastText"),
//   					Default: jsii.String("default"),
//   				},
//   			},
//   			Typography: &WorkspaceThemeTypographyProperty{
//   				FontFamily: &FontFamilyProperty{
//   					Default: jsii.String("default"),
//   				},
//   			},
//   		},
//   		Light: &WorkspaceThemeConfigProperty{
//   			Palette: &WorkspaceThemePaletteProperty{
//   				Canvas: &PaletteCanvasProperty{
//   					ActiveBackground: jsii.String("activeBackground"),
//   					ContainerBackground: jsii.String("containerBackground"),
//   					PageBackground: jsii.String("pageBackground"),
//   				},
//   				Header: &PaletteHeaderProperty{
//   					Background: jsii.String("background"),
//   					InvertActionsColors: jsii.Boolean(false),
//   					Text: jsii.String("text"),
//   					TextHover: jsii.String("textHover"),
//   				},
//   				Navigation: &PaletteNavigationProperty{
//   					Background: jsii.String("background"),
//   					InvertActionsColors: jsii.Boolean(false),
//   					Text: jsii.String("text"),
//   					TextActive: jsii.String("textActive"),
//   					TextBackgroundActive: jsii.String("textBackgroundActive"),
//   					TextBackgroundHover: jsii.String("textBackgroundHover"),
//   					TextHover: jsii.String("textHover"),
//   				},
//   				Primary: &PalettePrimaryProperty{
//   					Active: jsii.String("active"),
//   					ContrastText: jsii.String("contrastText"),
//   					Default: jsii.String("default"),
//   				},
//   			},
//   			Typography: &WorkspaceThemeTypographyProperty{
//   				FontFamily: &FontFamilyProperty{
//   					Default: jsii.String("default"),
//   				},
//   			},
//   		},
//   	},
//   	Title: jsii.String("title"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html
//
type CfnWorkspaceMixinProps struct {
	// The resource ARNs associated with the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-associations
	//
	Associations *[]*string `field:"optional" json:"associations" yaml:"associations"`
	// The description of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The media items for the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-media
	//
	Media interface{} `field:"optional" json:"media" yaml:"media"`
	// The name of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The pages associated with the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-pages
	//
	Pages interface{} `field:"optional" json:"pages" yaml:"pages"`
	// The tags used to organize, track, or control access for the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The theme configuration for the workspace, including colors and styling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-theme
	//
	Theme interface{} `field:"optional" json:"theme" yaml:"theme"`
	// The title displayed for the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Controls who can access the workspace.
	//
	// Valid values are: `ALL` (all users), `ASSIGNED` (only assigned users and routing profiles), and `NONE` (not visible).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html#cfn-connect-workspace-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

