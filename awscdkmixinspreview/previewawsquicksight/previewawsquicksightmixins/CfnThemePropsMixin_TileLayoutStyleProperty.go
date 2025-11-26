package previewawsquicksightmixins


// The display options for the layout of tiles on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tileLayoutStyleProperty := &TileLayoutStyleProperty{
//   	Gutter: &GutterStyleProperty{
//   		Show: jsii.Boolean(false),
//   	},
//   	Margin: &MarginStyleProperty{
//   		Show: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilelayoutstyle.html
//
type CfnThemePropsMixin_TileLayoutStyleProperty struct {
	// The gutter settings that apply between tiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilelayoutstyle.html#cfn-quicksight-theme-tilelayoutstyle-gutter
	//
	Gutter interface{} `field:"optional" json:"gutter" yaml:"gutter"`
	// The margin settings that apply around the outside edge of sheets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilelayoutstyle.html#cfn-quicksight-theme-tilelayoutstyle-margin
	//
	Margin interface{} `field:"optional" json:"margin" yaml:"margin"`
}

