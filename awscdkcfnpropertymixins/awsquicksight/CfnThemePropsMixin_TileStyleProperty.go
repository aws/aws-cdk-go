package awsquicksight


// Display options related to tiles on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tileStyleProperty := &TileStyleProperty{
//   	BackgroundColor: jsii.String("backgroundColor"),
//   	Border: &BorderStyleProperty{
//   		Color: jsii.String("color"),
//   		Show: jsii.Boolean(false),
//   		Width: jsii.String("width"),
//   	},
//   	BorderRadius: jsii.String("borderRadius"),
//   	Padding: jsii.String("padding"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilestyle.html
//
type CfnThemePropsMixin_TileStyleProperty struct {
	// String to encapsulate the most generic way Color can be formatted (words, hexStrings etc).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilestyle.html#cfn-quicksight-theme-tilestyle-backgroundcolor
	//
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The border around a tile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilestyle.html#cfn-quicksight-theme-tilestyle-border
	//
	Border interface{} `field:"optional" json:"border" yaml:"border"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilestyle.html#cfn-quicksight-theme-tilestyle-borderradius
	//
	BorderRadius *string `field:"optional" json:"borderRadius" yaml:"borderRadius"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilestyle.html#cfn-quicksight-theme-tilestyle-padding
	//
	Padding *string `field:"optional" json:"padding" yaml:"padding"`
}

