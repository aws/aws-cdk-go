package mixinsawsquicksight


// The theme display options for sheets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetStyleProperty := &SheetStyleProperty{
//   	Tile: &TileStyleProperty{
//   		Border: &BorderStyleProperty{
//   			Show: jsii.Boolean(false),
//   		},
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
type CfnThemePropsMixin_SheetStyleProperty struct {
	// The display options for tiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetstyle.html#cfn-quicksight-theme-sheetstyle-tile
	//
	Tile interface{} `field:"optional" json:"tile" yaml:"tile"`
	// The layout options for tiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetstyle.html#cfn-quicksight-theme-sheetstyle-tilelayout
	//
	TileLayout interface{} `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

