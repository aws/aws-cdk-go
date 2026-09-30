package awsquicksight


// The theme display options for sheets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetStyleProperty := &SheetStyleProperty{
//   	Background: &SheetBackgroundStyleProperty{
//   		Color: jsii.String("color"),
//   		Gradient: jsii.String("gradient"),
//   	},
//   	Tile: &TileStyleProperty{
//   		BackgroundColor: jsii.String("backgroundColor"),
//   		Border: &BorderStyleProperty{
//   			Color: jsii.String("color"),
//   			Show: jsii.Boolean(false),
//   			Width: jsii.String("width"),
//   		},
//   		BorderRadius: jsii.String("borderRadius"),
//   		Padding: jsii.String("padding"),
//   	},
//   	TileLayout: &TileLayoutStyleProperty{
//   		Gutter: &GutterStyleProperty{
//   			Show: jsii.Boolean(false),
//   		},
//   		Margin: &MarginStyleProperty{
//   			Show: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetstyle.html
//
type CfnTheme_SheetStyleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetstyle.html#cfn-quicksight-theme-sheetstyle-background
	//
	Background interface{} `field:"optional" json:"background" yaml:"background"`
	// The display options for tiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetstyle.html#cfn-quicksight-theme-sheetstyle-tile
	//
	Tile interface{} `field:"optional" json:"tile" yaml:"tile"`
	// The layout options for tiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetstyle.html#cfn-quicksight-theme-sheetstyle-tilelayout
	//
	TileLayout interface{} `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

