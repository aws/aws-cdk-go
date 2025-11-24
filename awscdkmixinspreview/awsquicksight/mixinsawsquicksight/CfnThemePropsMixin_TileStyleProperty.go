package mixinsawsquicksight


// Display options related to tiles on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tileStyleProperty := &TileStyleProperty{
//   	Border: &BorderStyleProperty{
//   		Show: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilestyle.html
//
type CfnThemePropsMixin_TileStyleProperty struct {
	// The border around a tile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilestyle.html#cfn-quicksight-theme-tilestyle-border
	//
	Border interface{} `field:"optional" json:"border" yaml:"border"`
}

