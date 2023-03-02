package awsquicksight


// Display options related to tiles on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tileStyleProperty := &tileStyleProperty{
//   	border: &borderStyleProperty{
//   		show: jsii.Boolean(false),
//   	},
//   }
//
type CfnTheme_TileStyleProperty struct {
	// The border around a tile.
	Border interface{} `field:"optional" json:"border" yaml:"border"`
}

