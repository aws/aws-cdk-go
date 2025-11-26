package previewawsquicksightmixins


// The display options for gutter spacing between tiles on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gutterStyleProperty := &GutterStyleProperty{
//   	Show: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-gutterstyle.html
//
type CfnThemePropsMixin_GutterStyleProperty struct {
	// This Boolean value controls whether to display a gutter space between sheet tiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-gutterstyle.html#cfn-quicksight-theme-gutterstyle-show
	//
	Show interface{} `field:"optional" json:"show" yaml:"show"`
}

