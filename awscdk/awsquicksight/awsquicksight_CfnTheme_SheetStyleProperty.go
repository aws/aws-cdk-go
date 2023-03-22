package awsquicksight


// The theme display options for sheets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetStyleProperty := &sheetStyleProperty{
//   	tile: &tileStyleProperty{
//   		border: &borderStyleProperty{
//   			show: jsii.Boolean(false),
//   		},
//   	},
//   	tileLayout: &tileLayoutStyleProperty{
//   		gutter: &gutterStyleProperty{
//   			show: jsii.Boolean(false),
//   		},
//   		margin: &marginStyleProperty{
//   			show: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnTheme_SheetStyleProperty struct {
	// The display options for tiles.
	Tile interface{} `field:"optional" json:"tile" yaml:"tile"`
	// The layout options for tiles.
	TileLayout interface{} `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

