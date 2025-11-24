package mixinsawsquicksight


// The display options for margins around the outside edge of sheets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   marginStyleProperty := &MarginStyleProperty{
//   	Show: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-marginstyle.html
//
type CfnThemePropsMixin_MarginStyleProperty struct {
	// This Boolean value controls whether to display sheet margins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-marginstyle.html#cfn-quicksight-theme-marginstyle-show
	//
	Show interface{} `field:"optional" json:"show" yaml:"show"`
}

