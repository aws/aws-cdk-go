package awsquicksight


// The display options for the layout of tiles on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tileLayoutStyleProperty := &tileLayoutStyleProperty{
//   	gutter: &gutterStyleProperty{
//   		show: jsii.Boolean(false),
//   	},
//   	margin: &marginStyleProperty{
//   		show: jsii.Boolean(false),
//   	},
//   }
//
type CfnTheme_TileLayoutStyleProperty struct {
	// The gutter settings that apply between tiles.
	Gutter interface{} `field:"optional" json:"gutter" yaml:"gutter"`
	// The margin settings that apply around the outside edge of sheets.
	Margin interface{} `field:"optional" json:"margin" yaml:"margin"`
}

